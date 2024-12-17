package services

import (
	"backend/models"
	"backend/repositories"
	"errors"
	"log"

	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

// Thêm các phương thức để xử lý logic nghiệp vụ
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.GetAllUsers()
}

// Đăng ký người dùng
func (s *UserService) RegisterUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	return s.UserRepo.RegisterUser(user)
}

// Đăng nhập người dùng
func (s *UserService) LoginUser(email, password string) (models.User, string, error) {
	user, err := s.UserRepo.LoginUser(email, password)
	if err != nil {
		return models.User{}, "", err
	}

	// Tạo JWT
	token, err := s.GenerateJWT(user.UserID)
	if err != nil {
		return models.User{}, "", err
	}

	return user, token, nil
}

// Thay đổi mật khẩu
func (s *UserService) ChangePassword(userID int, newPassword string) error {
	return s.UserRepo.ChangePassword(userID, newPassword)
}

// Gửi email đặt lại mật khẩu
func (s *UserService) SendResetPasswordEmail(email string) error {
	// Tạo một email mới
	m := gomail.NewMessage()
	m.SetHeader("From", "youJoyfulToys@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Reset Password")
	m.SetBody("text/plain", "Click here to reset your password: <link>")
	// Gửi email
	d := gomail.NewDialer("smtp.example.com", 587, "your_email@example.com", "your_email_password")
	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}
	return nil
}

// Thêm một cấu trúc để lưu thông tin JWT
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// Tạo JWT
func (s *UserService) GenerateJWT(userID int) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // Hết hạn sau 72 giờ
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your_secret_key")) // Thay thế bằng secret key của bạn
}

// Xóa tài khoản người dùng
func (s *UserService) DeleteUser(userID int) error {
	hasDependencies, err := s.UserRepo.UserHasDependencies(userID)
	if err != nil {
		return err
	}
	if hasDependencies {
		return errors.New("user has dependencies and cannot be deleted")
	}
	return s.UserRepo.DeleteUser(userID)
}

// Phương thức ResetPassword
func (s *UserService) ResetPassword(request models.ResetPasswordRequest) error {
	// Logic để đặt lại mật khẩu
	// Ví dụ: tìm người dùng theo email và cập nhật mật khẩu
	return nil
}

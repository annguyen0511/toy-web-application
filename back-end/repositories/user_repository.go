package repositories

import (
	"backend/models"
	"database/sql"
	"log"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Thêm các phương thức để tương tác với bảng Users
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT UserID, FullName, Email, PhoneNumber, Address, Role, CreatedAt FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		var createdAt string // Temporary variable to hold the CreatedAt value
		if err := rows.Scan(&user.UserID, &user.FullName, &user.Email, &user.PhoneNumber, &user.Address, &user.Role, &createdAt); err != nil {
			return nil, err
		}
		// Convert createdAt string to time.Time
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAt); err == nil {
			user.CreatedAt = parsedTime
		} else {
			println("Error parsing CreatedAt:", err.Error()) // Print error
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Đăng ký người dùng
func (r *UserRepository) RegisterUser(user models.User) error {
	_, err := r.DB.Exec("INSERT INTO Users (FullName, Email, PasswordHash, PhoneNumber, Address) VALUES (?, ?, ?, ?, ?)",
		user.FullName, user.Email, user.PasswordHash, user.PhoneNumber, user.Address)
	return err
}

// Đăng nhập người dùng
func (r *UserRepository) LoginUser(email, password string) (models.User, error) {
	log.Printf("Login attempt with Email: %s, Password: %s", email, password)
	var user models.User
	var createdAt []byte // Use []byte to handle the database value
	err := r.DB.QueryRow("SELECT UserID, FullName, Email, PhoneNumber, Address, Role, CreatedAt FROM Users WHERE Email = ? AND PasswordHash = ?", email, password).Scan(&user.UserID, &user.FullName, &user.Email, &user.PhoneNumber, &user.Address, &user.Role, &createdAt)
	if err != nil {
		return user, err
	}
	// Convert createdAt []byte to time.Time
	user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		log.Printf("Error parsing CreatedAt: %s", err.Error())
		return user, err
	}
	return user, nil
}

// Thay đổi mật khẩu
func (r *UserRepository) ChangePassword(userID int, newPassword string) error {
	_, err := r.DB.Exec("UPDATE Users SET PasswordHash = ? WHERE UserID = ?", newPassword, userID)
	return err
}

// Thêm phương thức để cập nhật mật khẩu
func (r *UserRepository) ResetPassword(email, newPassword string) error {
	_, err := r.DB.Exec("UPDATE Users SET PasswordHash = ? WHERE Email = ?", newPassword, email)
	return err
}

// Xóa tài khoản người dùng
func (r *UserRepository) DeleteUser(userID int) error {
	_, err := r.DB.Exec("DELETE FROM Users WHERE UserID = ?", userID)
	return err
}

// Kiểm tra xem người dùng có tồn tại trong các bảng khác không
func (r *UserRepository) UserHasDependencies(userID int) (bool, error) {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM Orders WHERE UserID = ?", userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

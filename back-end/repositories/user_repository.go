package repositories

import (
	"backend/models"
	"database/sql"
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
		if err := rows.Scan(&user.UserID, &user.FullName, &user.Email, &user.PhoneNumber, &user.Address, &user.Role, &user.CreatedAt); err != nil {
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
	var user models.User
	err := r.DB.QueryRow("SELECT UserID, FullName, Email, PhoneNumber, Address, Role, CreatedAt FROM Users WHERE Email = ? AND PasswordHash = ?", email, password).Scan(&user.UserID, &user.FullName, &user.Email, &user.PhoneNumber, &user.Address, &user.Role, &user.CreatedAt)
	return user, err
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

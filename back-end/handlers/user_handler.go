package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{UserService: service}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Đăng ký người dùng
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.UserService.RegisterUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to register user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Đăng nhập người dùng
func (h *UserHandler) LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user, token, err := h.UserService.LoginUser(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

// Thay đổi mật khẩu
func (h *UserHandler) ChangePassword(c *gin.Context) {
	var data struct {
		UserID      int    `json:"user_id"`
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.UserService.ChangePassword(data.UserID, data.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to change password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

// Quên mật khẩu
func (h *UserHandler) ResetPassword(c *gin.Context) {
	var resetRequest models.ResetPasswordRequest
	if err := c.ShouldBindJSON(&resetRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.UserService.ResetPassword(resetRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to reset password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

// Gửi email đặt lại mật khẩu
func (h *UserHandler) SendResetPasswordEmail(c *gin.Context) {
	var data struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.UserService.SendResetPasswordEmail(data.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to send reset password email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reset password email sent successfully"})
}

// Xóa tài khoản người dùng
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("user_id")           // Lấy user_id từ URL
	userIDInt, err := strconv.Atoi(userID) // Chuyển đổi từ string sang int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := h.UserService.DeleteUser(userIDInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

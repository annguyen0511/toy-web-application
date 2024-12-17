package routes

import (
	"backend/handlers"
	"backend/models"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, db *models.DB) {
	userRepo := repositories.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	api := r.Group("/api/users")
	{
		api.GET("/", userHandler.GetUsers)                                         // Lấy tất cả người dùng
		api.POST("/register", userHandler.RegisterUser)                            // Đăng ký người dùng
		api.POST("/login", userHandler.LoginUser)                                  // Đăng nhập người dùng
		api.POST("/change-password", userHandler.ChangePassword)                   // Thay đổi mật khẩu
		api.POST("/reset-password", userHandler.ResetPassword)                     // Quên mật khẩu
		api.POST("/send-reset-password-email", userHandler.SendResetPasswordEmail) // Gửi email đặt lại mật khẩu
		api.DELETE("/:user_id", userHandler.DeleteUser)                            // Xóa tài khoản người dùng
	}
}

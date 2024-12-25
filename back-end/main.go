package main

import (
	"backend/models"
	"backend/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Middleware để ghi lại thông tin yêu cầu
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL)
		c.Next() // Tiếp tục xử lý yêu cầu
	}
}

func main() {
	// Kết nối cơ sở dữ liệu
	models.Connect()

	// Tạo router và cấu hình các routes
	r := routes.SetupRouter(models.GlobalDB)

	// CORS cho tất cả các nguồn và phương thức
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                         // Cho phép tất cả các nguồn
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},             // Cập nhật phương thức
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"}, // Cập nhật header
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Thêm middleware Logger
	r.Use(Logger())

	// Lấy port từ biến môi trường
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Mặc định nếu không có biến môi trường
	}

	// Chạy server
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

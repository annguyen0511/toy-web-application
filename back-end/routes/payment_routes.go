package routes

import (
	"backend/handlers"
	"backend/models"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine, db *models.DB) {
	paymentRepo := repositories.NewPaymentRepository(db.DB)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	api := r.Group("/api/payments")
	{
		api.POST("/", paymentHandler.AddPayment)                 // Thêm thanh toán
		api.PUT("/", paymentHandler.UpdatePayment)               // Sửa thanh toán
		api.DELETE("/:payment_id", paymentHandler.DeletePayment) // Xóa thanh toán
		api.GET("/", paymentHandler.GetAllPayments)              // Lấy tất cả thanh toán
		api.GET("/:payment_id", paymentHandler.GetPaymentByID)   // Lấy thanh toán theo ID
	}
}

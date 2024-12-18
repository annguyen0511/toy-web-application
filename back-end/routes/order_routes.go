package routes

import (
	"backend/handlers"
	"backend/models"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, db *models.DB) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	api := r.Group("/api/orders")
	{
		api.POST("/", orderHandler.AddOrder)                                      // Thêm đơn hàng
		api.POST("/detail", orderHandler.AddOrderDetail)                          // Thêm chi tiết đơn hàng
		api.PUT("/detail", orderHandler.UpdateOrderDetail)                        // Sửa chi tiết đơn hàng
		api.DELETE("/detail/:order_detail_id", orderHandler.DeleteOrderDetail)    // Xóa chi tiết đơn hàng
		api.GET("/user/:user_id", orderHandler.GetOrdersByUserID)                 // Lấy đơn hàng theo UserID
		api.GET("/", orderHandler.GetAllOrders)                                   // Lấy tất cả đơn hàng
		api.GET("/detail/order/:order_id", orderHandler.GetOrderDetailsByOrderID) // Lấy chi tiết đơn hàng theo OrderID
		api.GET("/detail", orderHandler.GetAllOrderDetails)                       // Lấy tất cả chi tiết đơn hàng
	}
}

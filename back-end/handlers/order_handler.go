package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderService *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: service}
}

// Thêm đơn hàng
func (h *OrderHandler) AddOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.OrderService.AddOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add order"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Order added successfully"})
}

// Thêm chi tiết đơn hàng
func (h *OrderHandler) AddOrderDetail(c *gin.Context) {
	var orderDetail models.OrderDetail
	if err := c.ShouldBindJSON(&orderDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.OrderService.AddOrderDetail(orderDetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add order detail"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Order detail added successfully"})
}

// Sửa chi tiết đơn hàng
func (h *OrderHandler) UpdateOrderDetail(c *gin.Context) {
	var orderDetail models.OrderDetail
	if err := c.ShouldBindJSON(&orderDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.OrderService.UpdateOrderDetail(orderDetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update order detail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order detail updated successfully"})
}

// Xóa chi tiết đơn hàng
func (h *OrderHandler) DeleteOrderDetail(c *gin.Context) {
	orderDetailID := c.Param("order_detail_id")          // Lấy order_detail_id từ URL
	orderDetailIDInt, err := strconv.Atoi(orderDetailID) // Chuyển đổi từ string sang int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order detail ID"})
		return
	}
	if err := h.OrderService.DeleteOrderDetail(orderDetailIDInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete order detail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order detail deleted successfully"})
}

// Lấy tất cả đơn hàng theo UserID
func (h *OrderHandler) GetOrdersByUserID(c *gin.Context) {
	userID := c.Param("user_id")           // Lấy user_id từ URL
	userIDInt, err := strconv.Atoi(userID) // Chuyển đổi từ string sang int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	orders, err := h.OrderService.GetOrdersByUserID(userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// Lấy tất cả đơn hàng
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.OrderService.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// Lấy tất cả chi tiết đơn hàng theo OrderID
func (h *OrderHandler) GetOrderDetailsByOrderID(c *gin.Context) {
	orderID := c.Param("order_id")           // Lấy order_id từ URL
	orderIDInt, err := strconv.Atoi(orderID) // Chuyển đổi từ string sang int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}
	orderDetails, err := h.OrderService.GetOrderDetailsByOrderID(orderIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch order details"})
		return
	}
	c.JSON(http.StatusOK, orderDetails)
}

// Lấy tất cả chi tiết đơn hàng
func (h *OrderHandler) GetAllOrderDetails(c *gin.Context) {
	orderDetails, err := h.OrderService.GetAllOrderDetails()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch order details"})
		return
	}
	c.JSON(http.StatusOK, orderDetails)
}

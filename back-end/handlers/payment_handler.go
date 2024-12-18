package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	PaymentService *services.PaymentService
}

func NewPaymentHandler(service *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{PaymentService: service}
}

// Thêm thanh toán
func (h *PaymentHandler) AddPayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.PaymentService.AddPayment(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add payment"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Payment added successfully"})
}

// Sửa thanh toán
func (h *PaymentHandler) UpdatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.PaymentService.UpdatePayment(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update payment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Payment updated successfully"})
}

// Xóa thanh toán
func (h *PaymentHandler) DeletePayment(c *gin.Context) {
	paymentID := c.Param("payment_id")           // Lấy payment_id từ URL
	paymentIDInt, err := strconv.Atoi(paymentID) // Chuyển đổi từ string sang int
	if err != nil {
		// Xử lý lỗi
	}
	if err := h.PaymentService.DeletePayment(paymentIDInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete payment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}

// Lấy tất cả thanh toán
func (h *PaymentHandler) GetAllPayments(c *gin.Context) {
	payments, err := h.PaymentService.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch payments"})
		return
	}
	c.JSON(http.StatusOK, payments)
}

// Lấy thanh toán theo ID
func (h *PaymentHandler) GetPaymentByID(c *gin.Context) {
	paymentID := c.Param("payment_id")           // Lấy payment_id từ URL
	paymentIDInt, err := strconv.Atoi(paymentID) // Chuyển đổi từ string sang int
	if err != nil {
		// Xử lý lỗi
	}
	payment, err := h.PaymentService.GetPaymentByID(paymentIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch payment"})
		return
	}
	c.JSON(http.StatusOK, payment)
}

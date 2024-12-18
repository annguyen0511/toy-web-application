package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	CartService *services.CartService
}

func NewCartHandler(service *services.CartService) *CartHandler {
	return &CartHandler{CartService: service}
}

// Thêm giỏ hàng
func (h *CartHandler) AddCart(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.CartService.AddCart(cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add cart"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Cart added successfully"})
}

// Thêm chi tiết giỏ hàng
func (h *CartHandler) AddCartDetail(c *gin.Context) {
	var cartDetail models.CartDetail
	if err := c.ShouldBindJSON(&cartDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.CartService.AddCartDetail(cartDetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add cart detail"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Cart detail added successfully"})
}

// Sửa chi tiết giỏ hàng
func (h *CartHandler) UpdateCartDetail(c *gin.Context) {
	var cartDetail models.CartDetail
	if err := c.ShouldBindJSON(&cartDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.CartService.UpdateCartDetail(cartDetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update cart detail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart detail updated successfully"})
}

// Xóa chi tiết giỏ hàng
func (h *CartHandler) DeleteCartDetail(c *gin.Context) {
	cartDetailID := c.Param("cart_detail_id")          // Lấy cart_detail_id từ URL
	cartDetailIDInt, err := strconv.Atoi(cartDetailID) // Chuyển đổi từ string sang int
	if err != nil {
		// Xử lý lỗi
	}
	if err := h.CartService.DeleteCartDetail(cartDetailIDInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete cart detail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart detail deleted successfully"})
}

// Lấy tất cả chi tiết giỏ hàng theo UserID
func (h *CartHandler) GetCartByUserID(c *gin.Context) {
	userID := c.Param("user_id")           // Lấy user_id từ URL
	userIDInt, err := strconv.Atoi(userID) // Chuyển đổi từ string sang int
	if err != nil {
		// Xử lý lỗi
	}
	cartDetails, err := h.CartService.GetCartByUserID(userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch cart details"})
		return
	}
	c.JSON(http.StatusOK, cartDetails)
}

// Lấy tất cả giỏ hàng
func (h *CartHandler) GetAllCarts(c *gin.Context) {
	carts, err := h.CartService.GetAllCarts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch carts"})
		return
	}
	c.JSON(http.StatusOK, carts)
}

// Xóa giỏ hàng
func (h *CartHandler) DeleteCart(c *gin.Context) {
	cartID := c.Param("cart_id")           // Lấy cart_id từ URL
	cartIDInt, err := strconv.Atoi(cartID) // Chuyển đổi từ string sang int
	if err != nil {
		// Xử lý lỗi
	}
	if err := h.CartService.DeleteCart(cartIDInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete cart"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart deleted successfully"})
}

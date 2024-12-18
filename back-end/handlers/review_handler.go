package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	ReviewService *services.ReviewService
}

func NewReviewHandler(service *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{ReviewService: service}
}

// Thêm đánh giá
func (h *ReviewHandler) AddReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.ReviewService.AddReview(review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add review"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Review added successfully"})
}

// Sửa đánh giá
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.ReviewService.UpdateReview(review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update review"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review updated successfully"})
}

// Xóa đánh giá
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	reviewID := c.Param("review_id")           // Lấy review_id từ URL
	reviewIDInt, err := strconv.Atoi(reviewID) // Chuyển đổi từ string sang int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}
	if err := h.ReviewService.DeleteReview(reviewIDInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete review"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}

// Lấy tất cả đánh giá
func (h *ReviewHandler) GetAllReviews(c *gin.Context) {
	reviews, err := h.ReviewService.GetAllReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch reviews"})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

// Lấy đánh giá theo sản phẩm
func (h *ReviewHandler) GetReviewsByProductID(c *gin.Context) {
	productID := c.Param("product_id")           // Lấy product_id từ URL
	productIDInt, err := strconv.Atoi(productID) // Chuyển đổi từ string sang int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	reviews, err := h.ReviewService.GetReviewsByProductID(productIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch reviews"})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

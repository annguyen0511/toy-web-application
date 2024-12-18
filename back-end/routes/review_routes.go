package routes

import (
	"backend/handlers"
	"backend/models"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func ReviewRoutes(r *gin.Engine, db *models.DB) {
	reviewRepo := repositories.NewReviewRepository(db.DB)
	reviewService := services.NewReviewService(reviewRepo)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	api := r.Group("/api/reviews")
	{
		api.POST("/", reviewHandler.AddReview)                               // Thêm đánh giá
		api.PUT("/", reviewHandler.UpdateReview)                             // Sửa đánh giá
		api.DELETE("/:review_id", reviewHandler.DeleteReview)                // Xóa đánh giá
		api.GET("/", reviewHandler.GetAllReviews)                            // Lấy tất cả đánh giá
		api.GET("/product/:product_id", reviewHandler.GetReviewsByProductID) // Lấy đánh giá theo ID sản phẩm
	}
}

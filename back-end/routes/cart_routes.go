package routes

import (
	"backend/handlers"
	"backend/models"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine, db *models.DB) {
	cartRepo := repositories.NewCartRepository(db.DB)
	cartService := services.NewCartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService)

	r.POST("/api/carts/", cartHandler.AddCart)                                  // Thêm giỏ hàng
	r.POST("/api/carts/detail", cartHandler.AddCartDetail)                      // Thêm chi tiết giỏ hàng
	r.PUT("/api/carts/detail", cartHandler.UpdateCartDetail)                    // Sửa chi tiết giỏ hàng
	r.DELETE("/api/carts/detail/:cart_detail_id", cartHandler.DeleteCartDetail) // Xóa chi tiết giỏ hàng
	r.GET("/api/carts/user/:user_id", cartHandler.GetCartByUserID)              // Lấy giỏ hàng theo UserID
	r.GET("/api/carts/", cartHandler.GetAllCarts)                               // Lấy tất cả giỏ hàng
	r.DELETE("/api/carts/:cart_id", cartHandler.DeleteCart)                     // Xóa giỏ hàng
}

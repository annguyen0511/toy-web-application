package routes

import (
	"backend/handlers"
	"backend/models"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, db *models.DB) {
	productRepo := repositories.NewProductRepository(db.DB)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	api := r.Group("/api/products")
	{
		api.POST("/", productHandler.AddProduct)                                // Thêm sản phẩm
		api.PUT("/", productHandler.UpdateProduct)                              // Sửa sản phẩm
		api.DELETE("/:product_id", productHandler.DeleteProduct)                // Xóa sản phẩm
		api.GET("/", productHandler.GetAllProducts)                             // Lấy tất cả sản phẩm
		api.GET("/:product_id", productHandler.GetProductByID)                  // Lấy sản phẩm theo ID
		api.GET("/search", productHandler.SearchProducts)                       // Tìm kiếm sản phẩm
		api.GET("/category/:category_id", productHandler.GetProductsByCategory) // Lấy sản phẩm theo danh mục
	}
}

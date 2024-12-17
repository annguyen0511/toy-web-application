package routes

import (
	"backend/handlers"
	"backend/models"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine, db *models.DB) {
	categoryRepo := repositories.NewCategoryRepository(db.DB)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	api := r.Group("/api/categories")
	{
		api.POST("/", categoryHandler.AddCategory)                  // Thêm danh mục
		api.PUT("/", categoryHandler.UpdateCategory)                // Sửa danh mục
		api.DELETE("/:category_id", categoryHandler.DeleteCategory) // Xóa danh mục
		api.GET("/", categoryHandler.GetAllCategories)              // Lấy tất cả danh mục
		api.GET("/:category_id", categoryHandler.GetCategoryByID)   // Lấy danh mục theo ID
		api.GET("/search", categoryHandler.SearchCategories)        // Tìm kiếm danh mục
	}
}

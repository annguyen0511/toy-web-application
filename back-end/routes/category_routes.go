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

	r.POST("/api/categories/", categoryHandler.AddCategory)                      // Thêm danh mục
	r.PUT("/api/categories/update/:category_id", categoryHandler.UpdateCategory) // Sửa danh mục
	r.DELETE("/api/categories/:category_id", categoryHandler.DeleteCategory)     // Xóa danh mục
	r.GET("/api/categories/", categoryHandler.GetAllCategories)                  // Lấy tất cả danh mục
	r.GET("/api/categories/:category_id", categoryHandler.GetCategoryByID)       // Lấy danh mục theo ID
	r.GET("/api/categories/search", categoryHandler.SearchCategories)            // Tìm kiếm danh mục
}

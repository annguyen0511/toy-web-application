package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryService *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService: service}
}

// Thêm danh mục
func (h *CategoryHandler) AddCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.CategoryService.AddCategory(category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add category"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Category added successfully"})
}

// Sửa danh mục
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	println("Đã vào đây") // Log data
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.CategoryService.UpdateCategory(category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update category", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

// Xóa danh mục
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryID := c.Param("category_id")           // Lấy category_id từ URL
	categoryIDInt, err := strconv.Atoi(categoryID) // Chuyển đổi từ string sang int
	if err != nil {
		// Xử lý lỗi
	}
	if err := h.CategoryService.DeleteCategory(categoryIDInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// Lấy tất cả danh mục
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.CategoryService.GetAllCategories()
	if err != nil {
		// Log the specific error
		println("Error fetching categories:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// Lấy danh mục theo ID
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	categoryID := c.Param("category_id")           // Lấy category_id từ URL
	categoryIDInt, err := strconv.Atoi(categoryID) // Chuyển đổi từ string sang int
	if err != nil {
		// Xử lý lỗi
	}
	category, err := h.CategoryService.GetCategoryByID(categoryIDInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

// Tìm kiếm danh mục
func (h *CategoryHandler) SearchCategories(c *gin.Context) {
	name := c.Query("name") // Lấy tên từ query
	categories, err := h.CategoryService.SearchCategories(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to search categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

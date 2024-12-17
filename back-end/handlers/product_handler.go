package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: service}
}

// Thêm sản phẩm
func (h *ProductHandler) AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.ProductService.AddProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add product"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully"})
}

// Sửa sản phẩm
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.ProductService.UpdateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// Xóa sản phẩm
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productID := c.Param("product_id")           // Lấy product_id từ URL
	productIDInt, err := strconv.Atoi(productID) // Chuyển đổi từ string sang int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	if err := h.ProductService.DeleteProduct(productIDInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// Lấy tất cả sản phẩm
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.ProductService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// Lấy sản phẩm theo ID
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	productID := c.Param("product_id")           // Lấy product_id từ URL
	productIDInt, err := strconv.Atoi(productID) // Chuyển đổi từ string sang int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	product, err := h.ProductService.GetProductByID(productIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch product"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Tìm kiếm sản phẩm
func (h *ProductHandler) SearchProducts(c *gin.Context) {
	name := c.Query("name") // Lấy tên từ query
	products, err := h.ProductService.SearchProducts(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to search products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// Lấy sản phẩm theo danh mục
func (h *ProductHandler) GetProductsByCategory(c *gin.Context) {
	categoryID := c.Param("category_id")           // Lấy category_id từ URL
	categoryIDInt, err := strconv.Atoi(categoryID) // Chuyển đổi từ string sang int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}
	products, err := h.ProductService.GetProductsByCategory(categoryIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

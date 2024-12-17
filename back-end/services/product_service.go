package services

import (
	"backend/models"
	"backend/repositories"
)

type ProductService struct {
	ProductRepo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: repo}
}

// Thêm sản phẩm
func (s *ProductService) AddProduct(product models.Product) error {
	return s.ProductRepo.AddProduct(product)
}

// Sửa sản phẩm
func (s *ProductService) UpdateProduct(product models.Product) error {
	return s.ProductRepo.UpdateProduct(product)
}

// Xóa sản phẩm
func (s *ProductService) DeleteProduct(productID int) error {
	return s.ProductRepo.DeleteProduct(productID)
}

// Lấy tất cả sản phẩm
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.ProductRepo.GetAllProducts()
}

// Lấy sản phẩm theo ID
func (s *ProductService) GetProductByID(productID int) (models.Product, error) {
	return s.ProductRepo.GetProductByID(productID)
}

// Tìm kiếm sản phẩm
func (s *ProductService) SearchProducts(name string) ([]models.Product, error) {
	return s.ProductRepo.SearchProducts(name)
}

// Lấy sản phẩm theo danh mục
func (s *ProductService) GetProductsByCategory(categoryID int) ([]models.Product, error) {
	return s.ProductRepo.GetProductsByCategory(categoryID)
}

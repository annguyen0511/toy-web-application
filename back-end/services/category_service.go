package services

import (
	"backend/models"
	"backend/repositories"
)

type CategoryService struct {
	CategoryRepo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepo: repo}
}

// Thêm danh mục
func (s *CategoryService) AddCategory(category models.Category) error {
	return s.CategoryRepo.AddCategory(category)
}

// Sửa danh mục
func (s *CategoryService) UpdateCategory(category models.Category) error {
	return s.CategoryRepo.UpdateCategory(category)
}

// Xóa danh mục
func (s *CategoryService) DeleteCategory(categoryID int) error {
	return s.CategoryRepo.DeleteCategory(categoryID)
}

// Lấy tất cả danh mục
func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.CategoryRepo.GetAllCategories()
}

// Lấy danh mục theo ID
func (s *CategoryService) GetCategoryByID(categoryID int) (models.Category, error) {
	return s.CategoryRepo.GetCategoryByID(categoryID)
}

// Tìm kiếm danh mục
func (s *CategoryService) SearchCategories(name string) ([]models.Category, error) {
	return s.CategoryRepo.SearchCategories(name)
}

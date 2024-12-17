package repositories

import (
	"backend/models"
	"database/sql"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// Thêm danh mục
func (r *CategoryRepository) AddCategory(category models.Category) error {
	_, err := r.db.Exec("INSERT INTO Categories (CategoryName, Description) VALUES (?, ?)",
		category.CategoryName, category.Description)
	return err
}

// Sửa danh mục
func (r *CategoryRepository) UpdateCategory(category models.Category) error {
	_, err := r.db.Exec("UPDATE Categories SET CategoryName = ?, Description = ? WHERE CategoryID = ?",
		category.CategoryName, category.Description, category.CategoryID)
	return err
}

// Xóa danh mục
func (r *CategoryRepository) DeleteCategory(categoryID int) error {
	_, err := r.db.Exec("DELETE FROM Categories WHERE CategoryID = ?", categoryID)
	return err
}

// Lấy tất cả danh mục
func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	rows, err := r.db.Query("SELECT CategoryID, CategoryName, Description, CreatedAt FROM Categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.CategoryID, &category.CategoryName, &category.Description, &category.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

// Lấy danh mục theo ID
func (r *CategoryRepository) GetCategoryByID(categoryID int) (models.Category, error) {
	var category models.Category
	err := r.db.QueryRow("SELECT CategoryID, CategoryName, Description, CreatedAt FROM Categories WHERE CategoryID = ?", categoryID).Scan(&category.CategoryID, &category.CategoryName, &category.Description, &category.CreatedAt)
	return category, err
}

// Tìm kiếm danh mục
func (r *CategoryRepository) SearchCategories(name string) ([]models.Category, error) {
	rows, err := r.db.Query("SELECT CategoryID, CategoryName, Description, CreatedAt FROM Categories WHERE CategoryName LIKE ?", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.CategoryID, &category.CategoryName, &category.Description, &category.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

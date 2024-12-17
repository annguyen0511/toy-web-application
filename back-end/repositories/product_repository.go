package repositories

import (
	"backend/models"
	"database/sql"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Thêm sản phẩm
func (r *ProductRepository) AddProduct(product models.Product) error {
	_, err := r.db.Exec("INSERT INTO Products (ProductName, Description, Price, Stock, CategoryID, ImageURL) VALUES (?, ?, ?, ?, ?, ?)",
		product.ProductName, product.Description, product.Price, product.Stock, product.CategoryID, product.ImageURL)
	return err
}

// Sửa sản phẩm
func (r *ProductRepository) UpdateProduct(product models.Product) error {
	_, err := r.db.Exec("UPDATE Products SET ProductName = ?, Description = ?, Price = ?, Stock = ?, CategoryID = ?, ImageURL = ? WHERE ProductID = ?",
		product.ProductName, product.Description, product.Price, product.Stock, product.CategoryID, product.ImageURL, product.ProductID)
	return err
}

// Xóa sản phẩm
func (r *ProductRepository) DeleteProduct(productID int) error {
	_, err := r.db.Exec("DELETE FROM Products WHERE ProductID = ?", productID)
	return err
}

// Lấy tất cả sản phẩm
func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	rows, err := r.db.Query("SELECT ProductID, ProductName, Description, Price, Stock, CategoryID, ImageURL, CreatedAt FROM Products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Description, &product.Price, &product.Stock, &product.CategoryID, &product.ImageURL, &product.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// Lấy sản phẩm theo ID
func (r *ProductRepository) GetProductByID(productID int) (models.Product, error) {
	var product models.Product
	err := r.db.QueryRow("SELECT ProductID, ProductName, Description, Price, Stock, CategoryID, ImageURL, CreatedAt FROM Products WHERE ProductID = ?", productID).Scan(&product.ProductID, &product.ProductName, &product.Description, &product.Price, &product.Stock, &product.CategoryID, &product.ImageURL, &product.CreatedAt)
	return product, err
}

// Tìm kiếm sản phẩm
func (r *ProductRepository) SearchProducts(name string) ([]models.Product, error) {
	rows, err := r.db.Query("SELECT ProductID, ProductName, Description, Price, Stock, CategoryID, ImageURL, CreatedAt FROM Products WHERE ProductName LIKE ?", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Description, &product.Price, &product.Stock, &product.CategoryID, &product.ImageURL, &product.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// Lấy sản phẩm theo danh mục
func (r *ProductRepository) GetProductsByCategory(categoryID int) ([]models.Product, error) {
	rows, err := r.db.Query("SELECT ProductID, ProductName, Description, Price, Stock, CategoryID, ImageURL, CreatedAt FROM Products WHERE CategoryID = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Description, &product.Price, &product.Stock, &product.CategoryID, &product.ImageURL, &product.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

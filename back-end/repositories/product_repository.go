package repositories

import (
	"backend/models"
	"database/sql"
	"time"
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
		var createdAt string // Temporary variable to hold the CreatedAt value
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Description, &product.Price, &product.Stock, &product.CategoryID, &product.ImageURL, &createdAt); err != nil {
			return nil, err
		}
		// Convert createdAt string to time.Time
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAt); err == nil {
			product.CreatedAt = parsedTime
		} else {
			println("Error parsing CreatedAt:", err.Error()) // Print error
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// Lấy sản phẩm theo ID
func (r *ProductRepository) GetProductByID(productID int) (models.Product, error) {
	var product models.Product
	var createdAt string // Temporary variable to hold the CreatedAt value
	err := r.db.QueryRow("SELECT ProductID, ProductName, Description, Price, Stock, CategoryID, ImageURL, CreatedAt FROM Products WHERE ProductID = ?", productID).Scan(&product.ProductID, &product.ProductName, &product.Description, &product.Price, &product.Stock, &product.CategoryID, &product.ImageURL, &createdAt)
	// Convert createdAt string to time.Time
	if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAt); err == nil {
		product.CreatedAt = parsedTime
	} else {
		println("Error parsing CreatedAt:", err.Error()) // Print error
		return product, err
	}
	return product, err
}

// Tìm kiếm sản phẩm

func (r *ProductRepository) SearchProducts(name string) ([]models.Product, error) {
	rows, err := r.db.Query("SELECT ProductID, ProductName, Description, Price, Stock, CategoryID, ImageURL, CreatedAt FROM Products WHERE ProductName LIKE ?", "%"+name+"%")
	if err != nil {
		println("Error querying products:", err.Error()) // Print error
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		var createdAt string // Temporary variable to hold the CreatedAt value
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Description, &product.Price, &product.Stock, &product.CategoryID, &product.ImageURL, &createdAt); err != nil {
			println("Error scanning product:", err.Error()) // Print error
			return nil, err
		}
		// Convert createdAt string to time.Time
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAt); err == nil {
			product.CreatedAt = parsedTime
		} else {
			println("Error parsing CreatedAt:", err.Error()) // Print error
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
		var createdAt string // Temporary variable to hold the CreatedAt value
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Description, &product.Price, &product.Stock, &product.CategoryID, &product.ImageURL, &createdAt); err != nil {
			return nil, err
		}
		// Convert createdAt string to time.Time
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAt); err == nil {
			product.CreatedAt = parsedTime
		} else {
			println("Error parsing CreatedAt:", err.Error()) // Print error
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

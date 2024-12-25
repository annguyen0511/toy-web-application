package repositories

import (
	"backend/models"
	"database/sql"
	"time"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

// Thêm giỏ hàng
func (r *CartRepository) AddCart(cart models.Cart) error {
	_, err := r.db.Exec("INSERT INTO Carts (UserID, CreatedAt) VALUES (?, ?)",
		cart.UserID, cart.CreatedAt)
	return err
}

// Thêm chi tiết giỏ hàng
func (r *CartRepository) AddCartDetail(cartDetail models.CartDetail) error {
	_, err := r.db.Exec("INSERT INTO CartDetails (CartID, ProductID, Quantity, Price) VALUES (?, ?, ?, ?)",
		cartDetail.CartID, cartDetail.ProductID, cartDetail.Quantity, cartDetail.Price)
	return err
}

// Sửa chi tiết giỏ hàng
func (r *CartRepository) UpdateCartDetail(cartDetail models.CartDetail) error {
	_, err := r.db.Exec("UPDATE CartDetails SET Quantity = ?, Price = ? WHERE CartDetailID = ?",
		cartDetail.Quantity, cartDetail.Price, cartDetail.CartDetailID)
	return err
}

// Xóa chi tiết giỏ hàng
func (r *CartRepository) DeleteCartDetail(cartDetailID int) error {
	_, err := r.db.Exec("DELETE FROM CartDetails WHERE CartDetailID = ?", cartDetailID)
	return err
}

// Lấy tất cả chi tiết giỏ hàng theo UserID
func (r *CartRepository) GetCartByUserID(userID int) ([]models.CartDetail, error) {
	rows, err := r.db.Query("SELECT CartDetailID, CartID, ProductID, Quantity, Price FROM CartDetails WHERE CartID IN (SELECT CartID FROM Carts WHERE UserID = ?)", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartDetails []models.CartDetail
	for rows.Next() {
		var cartDetail models.CartDetail
		if err := rows.Scan(&cartDetail.CartDetailID, &cartDetail.CartID, &cartDetail.ProductID, &cartDetail.Quantity, &cartDetail.Price); err != nil {
			return nil, err
		}
		cartDetails = append(cartDetails, cartDetail)
	}
	return cartDetails, nil
}

// Lấy tất cả giỏ hàng
func (r *CartRepository) GetAllCarts() ([]models.Cart, error) {
	rows, err := r.db.Query("SELECT CartID, UserID, CreatedAt FROM Carts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []models.Cart
	for rows.Next() {
		var cart models.Cart
		var createdAt string // Temporary variable to hold the CreatedAt value
		if err := rows.Scan(&cart.CartID, &cart.UserID, &createdAt); err != nil {
			return nil, err
		}
		// Convert createdAt string to time.Time
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAt); err == nil {
			cart.CreatedAt = parsedTime
		} else {
			println("Error parsing CreatedAt:", err.Error()) // Print error
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

// Xóa giỏ hàng
func (r *CartRepository) DeleteCart(cartID int) error {
	_, err := r.db.Exec("DELETE FROM Carts WHERE CartID = ?", cartID)
	return err
}

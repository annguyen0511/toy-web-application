package repositories

import (
	"backend/models"
	"database/sql"
	"time"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Thêm đơn hàng
func (r *OrderRepository) AddOrder(order models.Order) error {
	_, err := r.db.Exec("INSERT INTO Orders (UserID, TotalPrice, CreatedAt) VALUES (?, ?, ?)",
		order.UserID, order.TotalPrice, order.CreatedAt)
	return err
}

// Thêm chi tiết đơn hàng
func (r *OrderRepository) AddOrderDetail(orderDetail models.OrderDetail) error {
	_, err := r.db.Exec("INSERT INTO OrderDetails (OrderID, ProductID, Quantity, Price) VALUES (?, ?, ?, ?)",
		orderDetail.OrderID, orderDetail.ProductID, orderDetail.Quantity, orderDetail.Price)
	return err
}

// Sửa chi tiết đơn hàng
func (r *OrderRepository) UpdateOrderDetail(orderDetail models.OrderDetail) error {
	_, err := r.db.Exec("UPDATE OrderDetails SET Quantity = ?, Price = ? WHERE OrderDetailID = ?",
		orderDetail.Quantity, orderDetail.Price, orderDetail.OrderDetailID)
	return err
}

// Xóa chi tiết đơn hàng
func (r *OrderRepository) DeleteOrderDetail(orderDetailID int) error {
	_, err := r.db.Exec("DELETE FROM OrderDetails WHERE OrderDetailID = ?", orderDetailID)
	return err
}

// Lấy tất cả đơn hàng theo UserID
func (r *OrderRepository) GetOrdersByUserID(userID int) ([]models.Order, error) {
	rows, err := r.db.Query("SELECT OrderID, UserID, TotalPrice, CreatedAt FROM Orders WHERE UserID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.OrderID, &order.UserID, &order.TotalPrice, &order.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// Lấy tất cả đơn hàng
func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	rows, err := r.db.Query("SELECT OrderID, UserID, TotalPrice, CreatedAt FROM Orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		var createdAt string // Temporary variable to hold the CreatedAt value
		if err := rows.Scan(&order.OrderID, &order.UserID, &order.TotalPrice, &createdAt); err != nil {
			return nil, err
		}
		// Convert createdAt string to time.Time
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAt); err == nil {
			order.CreatedAt = parsedTime
		} else {
			println("Error parsing CreatedAt:", err.Error()) // Print error
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// Lấy tất cả chi tiết đơn hàng theo OrderID
func (r *OrderRepository) GetOrderDetailsByOrderID(orderID int) ([]models.OrderDetail, error) {
	rows, err := r.db.Query("SELECT OrderDetailID, OrderID, ProductID, Quantity, Price FROM OrderDetails WHERE OrderID = ?", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderDetails []models.OrderDetail
	for rows.Next() {
		var orderDetail models.OrderDetail
		if err := rows.Scan(&orderDetail.OrderDetailID, &orderDetail.OrderID, &orderDetail.ProductID, &orderDetail.Quantity, &orderDetail.Price); err != nil {
			return nil, err
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	return orderDetails, nil
}

// Lấy tất cả chi tiết đơn hàng
func (r *OrderRepository) GetAllOrderDetails() ([]models.OrderDetail, error) {
	rows, err := r.db.Query("SELECT OrderDetailID, OrderID, ProductID, Quantity, Price FROM OrderDetails")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderDetails []models.OrderDetail
	for rows.Next() {
		var orderDetail models.OrderDetail
		if err := rows.Scan(&orderDetail.OrderDetailID, &orderDetail.OrderID, &orderDetail.ProductID, &orderDetail.Quantity, &orderDetail.Price); err != nil {
			return nil, err
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	return orderDetails, nil
}

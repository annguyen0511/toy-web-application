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
	_, err := r.db.Exec("INSERT INTO Orders (UserID, TotalAmount, CreatedAt) VALUES (?, ?, ?)",
		order.UserID, order.TotalAmount, order.CreatedAt)
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
	rows, err := r.db.Query("SELECT OrderID, UserID, OrderDate, TotalAmount, OrderStatus, ShippingAddress, CreatedAt FROM Orders WHERE UserID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		var orderDate, createdAt string // Temporary variables to hold the OrderDate and CreatedAt values
		if err := rows.Scan(&order.OrderID, &order.UserID, &orderDate, &order.TotalAmount, &order.OrderStatus, &order.ShippingAddress, &createdAt); err != nil {
			return nil, err
		}
		// Convert orderDate and createdAt strings to time.Time
		if parsedOrderDate, err := time.Parse("2006-01-02 15:04:05", orderDate); err == nil {
			order.OrderDate = parsedOrderDate
		} else {
			println("Error parsing OrderDate:", err.Error()) // Print error
			return nil, err
		}
		if parsedCreatedAt, err := time.Parse("2006-01-02 15:04:05", createdAt); err == nil {
			order.CreatedAt = parsedCreatedAt
		} else {
			println("Error parsing CreatedAt:", err.Error()) // Print error
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// Lấy tất cả đơn hàng
func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	println("đã vô") // Print "đã vô" at the start
	rows, err := r.db.Query("SELECT OrderID, UserID, OrderDate, TotalAmount, OrderStatus, ShippingAddress, CreatedAt FROM orders")
	if err != nil {
		println("Error executing query:", err.Error()) // Print error
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		var orderDate, createdAt string // Temporary variables to hold the OrderDate and CreatedAt values
		if err := rows.Scan(&order.OrderID, &order.UserID, &orderDate, &order.TotalAmount, &order.OrderStatus, &order.ShippingAddress, &createdAt); err != nil {
			println("Error scanning row:", err.Error()) // Print error
			return nil, err
		}
		// Convert orderDate and createdAt strings to time.Time
		if parsedOrderDate, err := time.Parse("2006-01-02 15:04:05", orderDate); err == nil {
			order.OrderDate = parsedOrderDate
		} else {
			println("Error parsing OrderDate:", err.Error()) // Print error
			return nil, err
		}
		if parsedCreatedAt, err := time.Parse("2006-01-02 15:04:05", createdAt); err == nil {
			order.CreatedAt = parsedCreatedAt
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

// Cập nhật trạng thái đơn hàng
func (r *OrderRepository) UpdateOrderStatus(orderID int, status string) error {
	_, err := r.db.Exec("UPDATE Orders SET OrderStatus = ? WHERE OrderID = ?", status, orderID)
	return err
}

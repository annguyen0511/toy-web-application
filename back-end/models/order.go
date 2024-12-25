package models

import "time"

type Order struct {
	OrderID         int       `json:"order_id"`
	UserID          int       `json:"user_id"`
	OrderDate       time.Time `json:"order_date"`
	TotalAmount     float64   `json:"total_amount"`
	OrderStatus     string    `json:"order_status"`
	ShippingAddress string    `json:"shipping_address"`
	CreatedAt       time.Time `json:"created_at"`
}

type OrderDetail struct {
	OrderDetailID int     `json:"order_detail_id"`
	OrderID       int     `json:"order_id"`
	ProductID     int     `json:"product_id"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
}

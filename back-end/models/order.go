package models

import "time"

type Order struct {
	OrderID    int       `json:"order_id"`
	UserID     int       `json:"user_id"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}

type OrderDetail struct {
	OrderDetailID int     `json:"order_detail_id"`
	OrderID       int     `json:"order_id"`
	ProductID     int     `json:"product_id"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
}

package models

import "time"

type Cart struct {
	CartID    int       `json:"cart_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CartDetail struct {
	CartDetailID int     `json:"cart_detail_id"`
	CartID       int     `json:"cart_id"`
	ProductID    int     `json:"product_id"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
}

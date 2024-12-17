package models

import "time"

type Review struct {
	ReviewID  int       `json:"review_id"`
	ProductID int       `json:"product_id"`
	UserID    int       `json:"user_id"`
	Rating    int       `json:"rating"` // Giả sử rating là số nguyên từ 1 đến 5
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

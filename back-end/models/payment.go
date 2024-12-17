package models

import "time"

type Payment struct {
	PaymentID   int       `json:"payment_id"`
	OrderID     int       `json:"order_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Status      string    `json:"status"` // e.g., "completed", "pending", "failed"
}

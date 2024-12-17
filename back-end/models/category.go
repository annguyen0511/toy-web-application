package models

import "time"

type Category struct {
	CategoryID   int       `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
}
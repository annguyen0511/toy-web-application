package models

import "time"

type User struct {
	UserID       int       `json:"user_id"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	PhoneNumber  string    `json:"phone_number"`
	Address      string    `json:"address"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

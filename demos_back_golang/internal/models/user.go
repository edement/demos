package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	IsTrainer bool      `json:"isTrainer"`
}

type RegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsTrainer bool   `json:"isTrainer"`
}

type UserResponse struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	IsTrainer bool   `json:"isTrainer"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

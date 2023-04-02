package payload

import "time"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token   string    `json:"token"`
	Exp     time.Time `json:"exp"`
	Message string    `json:"message"`
}

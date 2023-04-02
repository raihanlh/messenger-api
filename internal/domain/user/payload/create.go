package payload

import "gitlab.com/raihanlh/messenger-api/internal/model"

type CreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type CreateResponse struct {
	User    *model.User `json:"user"`
	Message string      `json:"message"`
}

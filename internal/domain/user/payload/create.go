package payload

import "gitlab.com/raihanlh/messenger-api/internal/model"

type CreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateResponse struct {
	User    *model.User `json:"user"`
	Message string      `json:"message"`
}

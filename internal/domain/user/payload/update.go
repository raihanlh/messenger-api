package payload

import "gitlab.com/raihanlh/messenger-api/internal/model"

type UpdateRequest struct {
	UserID   string `param:"id" json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateResponse struct {
	User    *model.User `json:"user"`
	Message string      `json:"message"`
}

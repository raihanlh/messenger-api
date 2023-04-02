package payload

import "gitlab.com/raihanlh/messenger-api/internal/model"

type GetByIdRequest struct {
	UserID string `param:"id"`
}

type GetByIdResponse struct {
	User    *model.User `json:"user"`
	Message string      `json:"message"`
}

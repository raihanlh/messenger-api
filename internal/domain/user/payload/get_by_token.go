package payload

import "gitlab.com/raihanlh/messenger-api/internal/model"

type GetByTokenRequest struct {
	Token string
}

type GetByTokenResponse struct {
	User    *model.User `json:"user"`
	Message string      `json:"message"`
}
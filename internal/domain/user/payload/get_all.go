package payload

import (
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"gitlab.com/raihanlh/messenger-api/pkg/pagination"
)

type GetAllRequest struct {
	pagination.Pagination
	Search string `query:"search"`
}

type GetAllResponse struct {
	*pagination.Pagination
	PaginatedData []*model.User `json:"paginatedData"`
	Message string `json:"message"`
}

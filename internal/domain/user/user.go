package user

import (
	"context"

	"github.com/labstack/echo/v4"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"gitlab.com/raihanlh/messenger-api/pkg/pagination"
)

type Repository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*model.User, error)
	GetAll(ctx context.Context, pgn *pagination.Pagination, req *payload.GetAllRequest) ([]*model.User, error)
}

type Usecase interface {
	Create(ctx context.Context, req *payload.CreateRequest) (*payload.CreateResponse, error)
	Update(ctx context.Context, req *payload.UpdateRequest) (*payload.UpdateResponse, error)
	Delete(ctx context.Context, req *payload.DeleteRequest) (*payload.DeleteResponse, error)
	GetById(ctx context.Context, req *payload.GetByIdRequest) (*payload.GetByIdResponse, error)
	GetAll(ctx context.Context, req *payload.GetAllRequest) (*payload.GetAllResponse, error)
}

type Handler interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	GetById(ctx echo.Context) error
	GetAll(ctx echo.Context) error
}

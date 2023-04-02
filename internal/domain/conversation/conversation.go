package conversation

import (
	"context"

	"github.com/labstack/echo/v4"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
)

type Repository interface {
	Create(ctx context.Context, conv *model.Conversation) (*model.Conversation, error)
}

type Usecase interface {
	Create(ctx context.Context, req *payload.CreateRequest) (*payload.CreateResponse, error)
}

type Handler interface {
	Create(ctx echo.Context) error
}

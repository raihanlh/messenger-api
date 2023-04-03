package message

import (
	"context"

	"github.com/labstack/echo/v4"
	"gitlab.com/raihanlh/messenger-api/internal/domain/message/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
)

type Repository interface {
	Create(ctx context.Context, message *model.Message) (*model.Message, error)
	GetAllByConversationId(ctx context.Context, conversationId string) ([]*model.Message, error)
}

type Usecase interface {
	Create(ctx context.Context, req *payload.CreateMessageRequest) (*payload.CreateMessageResponse, error)
	GetAllByConversationId(ctx context.Context, req *payload.GetMessagesByConvIdRequest) (*payload.GetMessagesByConvIdResponse, error)
}

type Handler interface {
	Create(ctx echo.Context) error
	GetByConversationId(ctx echo.Context) error
}

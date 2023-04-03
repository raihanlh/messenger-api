package conversation

import (
	"context"

	"github.com/labstack/echo/v4"
	"gitlab.com/raihanlh/messenger-api/internal/domain/conversation/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
)

type Repository interface {
	Create(ctx context.Context, conv *model.Conversation) (*model.Conversation, error)
	GetById(ctx context.Context, id string) (*model.Conversation, error)
	GetAllByUserId(ctx context.Context, userId string) ([]*model.Conversation, error)
	GetBySenderReceiverIds(ctx context.Context, senderId string, receiverId string) (*model.Conversation, error)
}

type Usecase interface {
	Create(ctx context.Context, req *payload.CreateConversationRequest) (*payload.CreateConversationResponse, error)
	GetById(ctx context.Context, req *payload.GetByIdConversationRequest) (*payload.GetByIdConversationResponse, error)
}

type Handler interface {
	// Create(ctx echo.Context) error
	GetById(ctx echo.Context) error
}

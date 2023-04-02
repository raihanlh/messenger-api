package conversation

import (
	"context"

	"gitlab.com/raihanlh/messenger-api/internal/domain/conversation/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
)

type Repository interface {
	Create(ctx context.Context, conv *model.Conversation) (*model.Conversation, error)
	GetAllByUserId(ctx context.Context, userId string) ([]*model.Conversation, error)
	GetBySenderReceiverIds(ctx context.Context, senderId string, receiverId string) (*model.Conversation, error)
}

type Usecase interface {
	Create(ctx context.Context, req *payload.CreateConversationRequest) (*payload.CreateConversationResponse, error)
}

type Handler interface {
	// Create(ctx echo.Context) error
}

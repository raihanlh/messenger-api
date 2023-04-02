package usecase

import (
	"context"

	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
	"gitlab.com/raihanlh/messenger-api/internal/domain/conversation"
	"gitlab.com/raihanlh/messenger-api/internal/domain/conversation/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"gitlab.com/raihanlh/messenger-api/pkg/logger"
	"go.uber.org/zap"
)

type ConversationUsecase struct {
	repositories *dependency.Repositories
}

func New(r *dependency.Repositories) conversation.Usecase {
	return &ConversationUsecase{
		repositories: r,
	}
}

func (u ConversationUsecase) Create(ctx context.Context, req *payload.CreateConversationRequest) (*payload.CreateConversationResponse, error) {
	log := logger.GetLogger(ctx)

	sender, err := u.repositories.User.GetById(ctx, req.SenderID)
	if err != nil {
		log.Error("Failed to get sender: ", zap.Error(err))
		return nil, err
	}
	receiver, err := u.repositories.User.GetById(ctx, req.ReceiverID)
	if err != nil {
		log.Error("Failed to get receiver: ", zap.Error(err))
		return nil, err
	}

	conv, err := u.repositories.Conversation.Create(ctx, &model.Conversation{
		SenderID:   sender.ID,
		ReceiverID: receiver.ID,
	})
	if err != nil {
		log.Error("Failed to create conversation: ", zap.Error(err))
		return nil, err
	}

	return &payload.CreateConversationResponse{
		Conversation: conv,
	}, nil
}

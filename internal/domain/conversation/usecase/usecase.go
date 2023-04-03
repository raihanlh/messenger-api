package usecase

import (
	"context"
	"errors"

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

func (u ConversationUsecase) GetById(ctx context.Context, req *payload.GetByIdConversationRequest) (*payload.GetByIdConversationResponse, error) {
	log := logger.GetLogger(ctx)

	conv, err := u.repositories.Conversation.GetById(ctx, req.ConversationID)
	if err != nil {
		log.Error("Failed to get conversation by id: ", zap.Error(err))
		return nil, err
	}
	if conv.SenderID == req.UserID {
		userWith, err := u.repositories.User.GetById(ctx, conv.ReceiverID)
		if err != nil {
			log.Error("Failed to get sender: ", zap.Error(err))
			return nil, err
		}
		return &payload.GetByIdConversationResponse{
			ConversationID: conv.ID,
			WithUser: &model.User{
				Model:    model.Model{ID: conv.ReceiverID},
				Name:     userWith.Name,
				PhotoURL: userWith.PhotoURL,
			},
		}, nil
	} else if conv.ReceiverID == req.UserID {
		userWith, err := u.repositories.User.GetById(ctx, conv.SenderID)
		if err != nil {
			log.Error("Failed to get sender: ", zap.Error(err))
			return nil, err
		}
		return &payload.GetByIdConversationResponse{
			ConversationID: conv.ID,
			WithUser: &model.User{
				Model:    model.Model{ID: conv.ReceiverID},
				Name:     userWith.Name,
				PhotoURL: userWith.PhotoURL,
			},
		}, nil
	}

	return nil, errors.New("unauthorized")
}

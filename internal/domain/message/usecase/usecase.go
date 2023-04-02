package usecase

import (
	"context"
	"time"

	"github.com/labstack/gommon/log"
	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
	"gitlab.com/raihanlh/messenger-api/internal/domain/message"
	"gitlab.com/raihanlh/messenger-api/internal/domain/message/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"go.uber.org/zap"
)

type MessageUsecase struct {
	repositories *dependency.Repositories
}

func New(r *dependency.Repositories) message.Usecase {
	return &MessageUsecase{
		repositories: r,
	}
}

func (u MessageUsecase) Create(ctx context.Context, req *payload.CreateMessageRequest) (*payload.CreateMessageResponse, error) {
	convo, err := u.repositories.Conversation.GetBySenderReceiverIds(ctx, req.SenderID, req.ReceiverID)
	if err != nil {
		log.Error("Failed to get conversation: ", zap.Error(err))
		return nil, err
	}
	if convo == nil {
		convo, err = u.repositories.Conversation.Create(ctx, &model.Conversation{
			SenderID:   req.SenderID,
			ReceiverID: req.ReceiverID,
		})
		if err != nil {
			log.Error("Failed to create conversation: ", zap.Error(err))
			return nil, err
		}
	}

	sender, err := u.repositories.User.GetById(ctx, req.SenderID)
	if err != nil {
		log.Error("Failed to create message: ", zap.Error(err))
		return nil, err
	}
	receiver, err := u.repositories.User.GetById(ctx, req.ReceiverID)
	if err != nil {
		log.Error("Failed to create message: ", zap.Error(err))
		return nil, err
	}

	msg, err := u.repositories.Message.Create(ctx, &model.Message{
		SendAt:         time.Now(),
		ConversationID: convo.ID,
		SenderID:       req.SenderID,
		MessageText:    req.Message,
	})
	if err != nil {
		log.Error("Failed to create message: ", zap.Error(err))
		return nil, err
	}
	return &payload.CreateMessageResponse{
		ID:          msg.ID,
		MessageText: msg.MessageText,
		Sender: &model.User{
			Model: model.Model{ID: sender.ID},
			Name:  sender.Name,
		},
		SendAt: msg.SendAt,
		ConversationResponse: payload.GetConversationResponse{
			ConversationID: convo.ID,
			WithUser: &model.User{
				Model:    model.Model{ID: receiver.ID},
				Name:     receiver.Name,
				PhotoURL: receiver.PhotoURL,
			},
		},
	}, nil
}

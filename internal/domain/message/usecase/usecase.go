package usecase

import (
	"context"
	"errors"
	"time"

	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
	"gitlab.com/raihanlh/messenger-api/internal/domain/message"
	"gitlab.com/raihanlh/messenger-api/internal/domain/message/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"gitlab.com/raihanlh/messenger-api/pkg/logger"
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
	log := logger.GetLogger(ctx)
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
		SentAt:         time.Now(),
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
		SentAt: msg.SentAt,
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

func (u MessageUsecase) GetAllByConversationId(ctx context.Context, req *payload.GetMessagesByConvIdRequest) (*payload.GetMessagesByConvIdResponse, error) {
	log := logger.GetLogger(ctx)
	convo, err := u.repositories.Conversation.GetById(ctx, req.ConversationID)
	if err != nil {
		log.Error("Failed to get conversation: ", zap.Error(err))
		return nil, err
	}
	if convo.SenderID != req.UserID && convo.ReceiverID != req.UserID {
		log.Error("Unauthorized: ", zap.Error(err))
		return nil, errors.New("unauthorized")
	}

	msgs, err := u.repositories.Message.GetAllByConversationId(ctx, req.ConversationID)
	if err != nil {
		log.Error("Failed get messages by conversation id: ", zap.Error(err))
		return nil, err
	}

	var res payload.GetMessagesByConvIdResponse = msgs
	return &res, nil
}

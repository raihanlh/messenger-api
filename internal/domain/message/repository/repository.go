package repository

import (
	"context"

	"gitlab.com/raihanlh/messenger-api/internal/domain/message"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MessageRepository struct {
	DB *gorm.DB
}

func New(gormDB *gorm.DB) message.Repository {
	return &MessageRepository{
		DB: gormDB,
	}
}

func (r MessageRepository) Create(ctx context.Context, message *model.Message) (*model.Message, error) {
	result := r.DB.WithContext(ctx).Model(message).Clauses(clause.OnConflict{DoNothing: true}).Create(&message)
	return message, result.Error
}

func (r MessageRepository) GetAllBySenderReceiverIds(ctx context.Context, senderId string, receiverId string) ([]*model.Message, error) {
	return nil, nil
}

func (r MessageRepository) GetAllByConversationId(ctx context.Context, email string) ([]*model.Message, error) {
	return nil, nil
}

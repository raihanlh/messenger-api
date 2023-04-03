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

func (r MessageRepository) GetAllByConversationId(ctx context.Context, conversationId string) ([]*model.Message, error) {
	var messages []*model.Message
	result := r.DB.WithContext(ctx).Preload("Sender", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name")
	}).Where("conversation_id = ?", conversationId).Select("messages.id", "messages.message_text", "messages.sent_at", "messages.sender_id").Find(&messages)
	return messages, result.Error
}

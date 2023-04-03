package repository

import (
	"context"

	"gitlab.com/raihanlh/messenger-api/internal/constant"
	"gitlab.com/raihanlh/messenger-api/internal/domain/conversation"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ConversationRepository struct {
	DB *gorm.DB
}

func New(gormDB *gorm.DB) conversation.Repository {
	return &ConversationRepository{
		DB: gormDB,
	}
}

func (r ConversationRepository) Create(ctx context.Context, conv *model.Conversation) (*model.Conversation, error) {
	result := r.DB.WithContext(ctx).Model(conv).Clauses(clause.OnConflict{DoNothing: true}).Create(&conv)
	return conv, result.Error
}

func (r ConversationRepository) GetById(ctx context.Context, id string) (*model.Conversation, error) {
	var conv *model.Conversation
	result := r.DB.WithContext(ctx).Table(constant.ConversationTable).Where("id = ?", id).First(&conv)
	return conv, result.Error
}

func (r ConversationRepository) GetAllByUserId(ctx context.Context, userId string) ([]*model.Conversation, error) {

	return nil, nil
}

func (r ConversationRepository) GetBySenderReceiverIds(ctx context.Context, senderId string, receiverId string) (*model.Conversation, error) {
	var conv *model.Conversation

	result := r.DB.WithContext(ctx).Table(constant.ConversationTable).Where("sender_id = ? AND receiver_id = ?", senderId, receiverId).
		Or("sender_id = ? AND receiver_id = ?", receiverId, senderId).Limit(1).Find(&conv)

	if result.RowsAffected > 0 {
		return conv, result.Error
	} else {
		return nil, result.Error
	}
}

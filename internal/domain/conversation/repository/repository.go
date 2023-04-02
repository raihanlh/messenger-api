package repository

import (
	"context"

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
	result := r.DB.WithContext(ctx).Model(conv).Clauses(clause.OnConflict{DoNothing: true}).Create(conv)
	return conv, result.Error
}

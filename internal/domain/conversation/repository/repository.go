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

// SELECT product_id, MAX(sale_date) as latest_sale_date, sale_amount
// FROM sales
// GROUP BY product_id
func (r ConversationRepository) GetAllByUserId(ctx context.Context, userId string) ([]*model.Conversation, error) {
	// convs := make([]*model.Conversation, 0)
	// var rows []map[string]interface{}
	// result := r.DB.WithContext(ctx).Select("c.id", "m.message_text as msg_text", "MAX(m.created_at) as latest_message_at").
	// 	Table(constant.ConversationTable+"AS c").Joins("JOIN user_participants p ON p.conversation_id = c.id").
	// 	Joins("JOIN messages m ON m.conversation_id = c.id").Where("p.user_id = ?", userId).Distinct().Group("c.id").
	// 	Order("latest_message_at DESC").Scan(&rows)

	// log.Println(rows)
	// for _, row := range rows {
	// 	conv := &model.Conversation{
	// 		Model: model.Model{
	// 			ID: row["id"].(string),
	// 		},
	// 		LatestMessage: model.Message{
	// 			MessageText: row["msg_text"].(string),
	// 		},
	// 	}
	// 	convs = append(convs, conv)
	// }

	// return convs, result.Error
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

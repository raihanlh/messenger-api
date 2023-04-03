package model

import (
	"time"

	"gitlab.com/raihanlh/messenger-api/internal/constant"
)

type Message struct {
	Model          `swaggerignore:"true"`
	SentAt         time.Time     `json:"sent_at" gorm:"autoCreateTime"`
	ConversationID string        `json:"conversationId,omitempty"`
	SenderID       string        `json:"-"`
	MessageText    string        `json:"message,omitempty"`
	Conversation   *Conversation `gorm:"foreignKey:ConversationID" json:"-"`
	Sender         *User         `gorm:"foreignKey:SenderID" json:"sender"`
	IsRead         bool          `gorm:"default:false" json:"-"`
}

// Table name for gorm
func (u *Message) Table() string {
	return constant.MessageTable
}

package model

import (
	"time"

	"gitlab.com/raihanlh/messenger-api/internal/constant"
)

type Message struct {
	Model          `swaggerignore:"true"`
	SendAt         time.Time     `json:"send_at" gorm:"autoCreateTime"`
	ConversationID string        `json:"conversationId"`
	SenderID       string        `json:"senderId"`
	MessageText    string        `json:"messageText"`
	Conversation   *Conversation `gorm:"foreignKey:ConversationID" json:"-"`
	Sender         *User         `gorm:"foreignKey:SenderID" json:"-"`
}

// Table name for gorm
func (u *Message) Table() string {
	return constant.MessageTable
}

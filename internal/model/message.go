package model

import "gitlab.com/raihanlh/messenger-api/internal/constant"

type Message struct {
	Model          `swaggerignore:"true"`
	ConversationID string        `json:"conversationId"`
	SenderID       string        `json:"senderId"`
	ReceiverID     string        `json:"receiverId"`
	MessageText    string        `json:"messageText"`
	Conversation   *Conversation `gorm:"foreignKey:ConversationID" json:"-"`
	Sender         *User         `gorm:"foreignKey:SenderID" json:"-"`
	Receiver       *User         `gorm:"foreignKey:ReceiverID" json:"-"`
}

// Table name for gorm
func (u *Message) Table() string {
	return constant.MessageTable
}

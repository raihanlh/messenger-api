package model

import "gitlab.com/raihanlh/messenger-api/internal/constant"

type Conversation struct {
	Model      `swaggerignore:"true"`
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Sender     *User  `gorm:"foreignKey:SenderID" json:"-"`
	Receiver   *User  `gorm:"foreignKey:ReceiverID" json:"-"`
}

// Table name for gorm
func (u *Conversation) Table() string {
	return constant.ConversationTable
}

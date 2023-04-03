package payload

import (
	"time"

	"gitlab.com/raihanlh/messenger-api/internal/model"
)

type CreateMessageRequest struct {
	Message    string `json:"message"`
	SenderID   string `json:"-"`
	ReceiverID string `json:"user_id"`
}

type CreateMessageResponse struct {
	ID                   string                  `json:"id"` // Message ID
	MessageText          string                  `json:"message"`
	Sender               *model.User             `json:"sender"`
	SentAt               time.Time               `json:"sent_at"`
	ConversationResponse GetConversationResponse `json:"conversation"`
}

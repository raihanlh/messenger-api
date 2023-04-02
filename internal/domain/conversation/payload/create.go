package payload

import "gitlab.com/raihanlh/messenger-api/internal/model"

type CreateConversationRequest struct {
	SenderID   string `json:"senderId"`
	ReceiverID string `json:"receiverId"`
}

type CreateConversationResponse struct {
	Conversation *model.Conversation
}

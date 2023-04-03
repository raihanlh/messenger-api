package payload

import "gitlab.com/raihanlh/messenger-api/internal/model"

type GetAllByUserIdConvRequest struct {
	UserID string `json:"-"`
}
type GetAllByUserIdConvResponse []*GetAllByUserIdConv

type GetAllByUserIdConv struct {
	GetByIdConversationResponse
	LastMessage *model.Message `json:"last_message"`
	UnreadCount int64          `json:"unread_count"`
}

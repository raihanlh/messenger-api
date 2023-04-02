package payload

import "gitlab.com/raihanlh/messenger-api/internal/model"

type GetConversationResponse struct {
	ConversationID string      `json:"id"`
	WithUser       *model.User `json:"with_user"`
}

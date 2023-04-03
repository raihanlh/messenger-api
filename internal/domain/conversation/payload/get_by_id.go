package payload

import "gitlab.com/raihanlh/messenger-api/internal/model"

type GetByIdConversationRequest struct {
	ConversationID string `param:"convo_id"`
	UserID         string `json:"-"`
}

type GetByIdConversationResponse struct {
	ConversationID string      `json:"id"`
	WithUser       *model.User `json:"with_user"`
}

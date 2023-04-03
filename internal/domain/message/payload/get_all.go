package payload

import "gitlab.com/raihanlh/messenger-api/internal/model"

type GetMessagesByConvIdRequest struct {
	ConversationID string `param:"convo_id"`
	UserID         string `json:"-"`
}

type GetMessagesByConvIdResponse []*model.Message

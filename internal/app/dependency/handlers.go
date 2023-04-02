package dependency

import (
	"gitlab.com/raihanlh/messenger-api/internal/domain/conversation"
	"gitlab.com/raihanlh/messenger-api/internal/domain/message"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user"
	"gitlab.com/raihanlh/messenger-api/internal/health"
)

type Handlers struct {
	User         user.Handler
	Health       health.Handler
	Message      message.Handler
	Conversation conversation.Handler
}

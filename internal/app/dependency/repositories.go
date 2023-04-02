package dependency

import (
	"gitlab.com/raihanlh/messenger-api/internal/domain/conversation"
	"gitlab.com/raihanlh/messenger-api/internal/domain/message"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user"
)

// Add repositories here
type Repositories struct {
	User         user.Repository
	Message      message.Repository
	Conversation conversation.Repository
}

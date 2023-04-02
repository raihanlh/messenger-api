package dependency

import (
	"gitlab.com/raihanlh/messenger-api/internal/domain/conversation"
	"gitlab.com/raihanlh/messenger-api/internal/domain/message"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user"
)

// Add usecases here
type Usecases struct {
	User         user.Usecase
	Message      message.Usecase
	Conversation conversation.Usecase
}

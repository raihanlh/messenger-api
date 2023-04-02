package handler

import (
	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
	"gitlab.com/raihanlh/messenger-api/internal/domain/conversation"
)

type ConversationHandler struct {
	usecases *dependency.Usecases
}

func New(u *dependency.Usecases) conversation.Handler {
	return &ConversationHandler{
		usecases: u,
	}
}

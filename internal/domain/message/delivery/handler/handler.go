package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	apiPayload "gitlab.com/raihanlh/messenger-api/api/payload"
	http_error "gitlab.com/raihanlh/messenger-api/api/payload/http-error"
	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
	"gitlab.com/raihanlh/messenger-api/internal/domain/message"
	"gitlab.com/raihanlh/messenger-api/internal/domain/message/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
)

type MessageHandler struct {
	usecases *dependency.Usecases
}

func New(u *dependency.Usecases) message.Handler {
	return &MessageHandler{
		usecases: u,
	}
}

// CreateNewMessage godoc
// @Summary Create New Message
// @Description create message from request body
// @Tags Message
// @Accept application/json
// @Param body body payload.CreateMessageRequest true "Create User"
// @Produce json
// @Success 201 {object} object{status=string,data=payload.CreateMessageResponse}
// @Router /api/v1/messages [post]
func (h MessageHandler) Create(ctx echo.Context) error {
	var body payload.CreateMessageRequest

	if err := ctx.Bind(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(errCustom.HTTPCode, errCustom.HttpResponseError())
	}

	// Validate incoming data
	user := ctx.Get("user").(*model.User)
	body.SenderID = user.ID
	if err := ctx.Validate(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(http.StatusBadRequest, errCustom)
	}

	// Check if message is empty
	if body.Message == "" {
		return ctx.JSON(http.StatusUnprocessableEntity, "message can't be empty")
	}

	// Pass body to usecase
	data, err := h.usecases.Message.Create(ctx.Request().Context(), &body)
	if err != nil {
		httpErr, ok := err.(*http_error.Error)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, http_error.InternalServerError(fmt.Sprintf("Failed to create user: %s", httpErr.Error())))
		}
		return ctx.JSON(httpErr.HTTPCode, httpErr.HttpResponseError())
	}

	res := new(apiPayload.BaseResponse)
	res.AddHTTPCode(http.StatusCreated).AddData(data)
	return ctx.JSON(res.HTTPCode, res)
}

// GetMessageByConversationId godoc
// @Summary Get Message By Conversation Id
// @Description get message by conversation id
// @Tags Message
// @Accept application/json
// @Param convo_id path string true "Conversation ID"
// @Produce json
// @Success 200 {object} object{status=string,data=payload.GetMessagesByConvIdResponse}
// @Router /api/v1/conversations/{convo_id}/messages [get]
func (h MessageHandler) GetByConversationId(ctx echo.Context) error {
	var body payload.GetMessagesByConvIdRequest

	if err := ctx.Bind(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(errCustom.HTTPCode, errCustom.HttpResponseError())
	}

	// Validate incoming data
	if err := ctx.Validate(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(http.StatusBadRequest, errCustom)
	}

	// Pass body to usecase
	user := ctx.Get("user").(*model.User)
	body.UserID = user.ID
	data, err := h.usecases.Message.GetAllByConversationId(ctx.Request().Context(), &body)
	if err != nil {
		if err.Error() == "unauthorized" {
			return ctx.JSON(http.StatusForbidden, "forbidden")
		}
		if err.Error() == "not found" {
			return ctx.JSON(http.StatusNotFound, "not found")
		}
		httpErr, ok := err.(*http_error.Error)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, http_error.InternalServerError(fmt.Sprintf("Failed to get user by id: %s", httpErr.Error())))
		}
		return ctx.JSON(httpErr.HTTPCode, httpErr.HttpResponseError())
	}
	var msgs []*model.Message = *data

	res := new(apiPayload.BaseResponse)
	res.AddHTTPCode(http.StatusOK).AddData(msgs)
	return ctx.JSON(res.HTTPCode, res)
}

package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	apiPayload "gitlab.com/raihanlh/messenger-api/api/payload"
	http_error "gitlab.com/raihanlh/messenger-api/api/payload/http-error"
	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
	"gitlab.com/raihanlh/messenger-api/internal/domain/conversation"
	"gitlab.com/raihanlh/messenger-api/internal/domain/conversation/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
)

type ConversationHandler struct {
	usecases *dependency.Usecases
}

func New(u *dependency.Usecases) conversation.Handler {
	return &ConversationHandler{
		usecases: u,
	}
}

// GetConversationById godoc
// @Summary Get Conversation By Id
// @Description get conversation by id
// @Tags Conversation
// @Accept application/json
// @Param convo_id path string true "Conversation ID"
// @Produce json
// @Success 200 {object} object{status=string,data=payload.GetByIdConversationResponse}
// @Router /api/v1/conversations/{convo_id} [get]
func (h ConversationHandler) GetById(ctx echo.Context) error {
	var body payload.GetByIdConversationRequest

	if err := ctx.Bind(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(errCustom.HTTPCode, errCustom.HttpResponseError())
	}

	log.Printf("%+v", body)
	// Validate incoming data
	if err := ctx.Validate(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(http.StatusBadRequest, errCustom)
	}

	// Pass body to usecase
	user := ctx.Get("user").(*model.User)
	body.UserID = user.ID
	data, err := h.usecases.Conversation.GetById(ctx.Request().Context(), &body)
	if err != nil {
		if err.Error() == "not found" {
			return ctx.JSON(http.StatusNotFound, http_error.RecordNotFound("conversation"))
		}
		if err.Error() == "unauthorized" {
			return ctx.JSON(http.StatusNotFound, http_error.Unauthorized("unathorized"))
		}
		httpErr, ok := err.(*http_error.Error)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, http_error.InternalServerError(fmt.Sprintf("Failed to get user by id: %s", httpErr.Error())))
		}
		return ctx.JSON(httpErr.HTTPCode, httpErr.HttpResponseError())
	}

	res := new(apiPayload.BaseResponse)
	res.AddHTTPCode(http.StatusOK).AddStatus(apiPayload.StatusOK).AddData(data)
	return ctx.JSON(res.HTTPCode, res)
}

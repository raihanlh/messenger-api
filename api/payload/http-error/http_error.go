package http_error

import (
	"fmt"
	"net/http"

	payload "gitlab.com/raihanlh/messenger-api/api/payload"
)

type Error struct {
	HTTPCode  int
	ErrorCode string
	Messages  interface{}
}

func (e Error) Error() string {
	return e.Messages.(string)
}

func CustomError(h int, ec string, m interface{}) *Error {
	return &Error{
		HTTPCode:  h,
		ErrorCode: ec,
		Messages:  m,
	}
}

func RecordNotFound(item string) *Error {
	httpCode := http.StatusNotFound
	errorCode := RecordNotFoundCode
	message := fmt.Sprintf("%s not found", item)
	return CustomError(httpCode, errorCode, message)
}

func BadRequest(error error) *Error {
	httpCode := http.StatusBadRequest
	errorCode := BadRequestCode
	message := error.Error()
	return CustomError(httpCode, errorCode, message)
}

func Unauthorized(msg string) *Error {
	httpCode := http.StatusUnauthorized
	errorCode := UnauthorizedCode
	message := msg
	return CustomError(httpCode, errorCode, message)
}

func Forbidden(msg string) *Error {
	httpCode := http.StatusForbidden
	errorCode := ForbiddenCode
	message := msg
	return CustomError(httpCode, errorCode, message)
}

func InternalServerError(msg string) *Error {
	httpCode := http.StatusInternalServerError
	errorCode := InternalServerErrorCode
	message := msg
	return CustomError(httpCode, errorCode, message)
}

// Build Error response
func (e Error) HttpResponseError() *payload.BaseResponse {
	httpResp := new(payload.BaseResponse)
	return httpResp.AddHTTPCode(e.HTTPCode).
		AddErrorCode(e.ErrorCode).
		AddMessages(e.Messages).
		AddStatus(payload.StatusError)
}

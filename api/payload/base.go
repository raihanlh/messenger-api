package payloadV1

const (
	StatusOK    = "OK"
	StatusError = "ERROR"
)

type BaseResponse struct {
	HTTPCode  int         `json:"-"`
	Status    string      `json:"status,omitempty"`
	ErrorCode string      `json:"errorCode,omitempty"`
	Messages  interface{} `json:"messages,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

type Response interface {
	AddStatus(status string) *BaseResponse
	AddErrorCode(errorCode string) *BaseResponse
	AddMessages(message interface{}) *BaseResponse
	AddData(data interface{}) *BaseResponse
	AddHTTPCode(httpCode int) *BaseResponse
}

func NewResponse() Response {
	return &BaseResponse{}
}

func (b *BaseResponse) AddStatus(s string) *BaseResponse {
	b.Status = s
	return b
}

func (b *BaseResponse) AddErrorCode(e string) *BaseResponse {
	b.ErrorCode = e
	return b
}

func (b *BaseResponse) AddMessages(i interface{}) *BaseResponse {
	b.Messages = i
	return b
}

func (b *BaseResponse) AddData(d interface{}) *BaseResponse {
	b.Data = d
	return b
}

func (b *BaseResponse) AddHTTPCode(h int) *BaseResponse {
	b.HTTPCode = h
	return b
}

package payload

type DeleteRequest struct {
	UserID string `param:"id"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

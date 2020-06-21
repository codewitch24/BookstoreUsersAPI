package errors

import "net/http"

const (
	BadRequest = "BAD_REQUEST"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   BadRequest,
	}
}

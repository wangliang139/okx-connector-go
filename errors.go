package okx_connector

import (
	"fmt"
)

// ApiError define API error when response status is 4xx or 5xx
type ApiError struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

// Error return error code and message
func (e ApiError) Error() string {
	return fmt.Sprintf("<ApiError> code=%s, msg=%s", e.Code, e.Message)
}

// IsApiError check if e is an API error
func IsApiError(e error) bool {
	_, ok := e.(*ApiError)
	return ok
}

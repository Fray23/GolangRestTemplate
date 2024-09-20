package custom_error

import (
	"fmt"
)

type ClientError struct {
	Code    int
	Message string
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

func NewClientError(code int, message string) *ClientError {
	return &ClientError{
		Code:    code,
		Message: message,
	}
}

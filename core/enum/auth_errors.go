package auth_errors

import "fmt"

type ClientError struct {
	Message  string
	Code     int
	HttpCode int
	extra    interface{}
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("Client Error: %s (Code: %d)", e.Message, e.Code)
}

// status codes
const (
	USER_ALREADY_REGISTERED = 10
)

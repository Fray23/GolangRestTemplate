package auth_errors

import (
	"errors"
)


var (
	ErrNotFound       = errors.New("resource not found")
	ErrInvalidInput   = errors.New("invalid input")
	ErrInternalServer = errors.New("internal server error")

	UserAlreadyRegistered = errors.New("user already registered1")
)

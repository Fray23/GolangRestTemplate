package auth_errors

import (
	"code/core/error"
)

var (
	UserAlreadyRegistered = custom_error.NewClientError(409, "user already registered")
)

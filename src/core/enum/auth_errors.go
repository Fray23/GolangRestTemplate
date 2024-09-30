package auth_errors

import (
	"code/core/error"
)

var (
	UserAlreadyRegistered = custom_error.NewClientError(409, "user already registered")
	UserNotFound = custom_error.NewClientError(410, "user not found")
	InvalidPassword = custom_error.NewClientError(411, "invalid password or login")
)

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth_err "code/core/enum"
)


func ErrorWrapper [T any](handlerFunc func(c *gin.Context) (T, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		response_dto, err := handlerFunc(c)
		if err != nil {
			switch err {
			case auth_err.UserAlreadyRegistered:
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error occurred"})
			}
			return
		}
		c.JSON(http.StatusOK, response_dto)
	}
}

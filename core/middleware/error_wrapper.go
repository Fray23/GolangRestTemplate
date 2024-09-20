package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"code/core/enum"
	"code/core/error"
)


func ErrorWrapper [T any](handlerFunc func(c *gin.Context) (T, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		response_dto, err := handlerFunc(c)
		if err != nil {
			switch err {
			case auth_errors.UserAlreadyRegistered:
				if err, ok := err.(*custom_error.ClientError); ok {
					c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "code": err.Code})
				} else {
					c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				}
			default:
			    fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error occurred"})
			}
			return
		}
		c.JSON(http.StatusOK, response_dto)
	}
}

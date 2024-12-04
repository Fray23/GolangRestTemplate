package middleware

import (
	"net/http"

	"code/core/security"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.Request.Header.Get("Authorization")
		if jwtToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		}
		_, err := security.DecodeJwtToken(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		}
	}
}

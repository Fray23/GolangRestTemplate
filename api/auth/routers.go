package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.GET("/login", LoginHandler)
	rg.POST("/signup", SignUpHandler)
}

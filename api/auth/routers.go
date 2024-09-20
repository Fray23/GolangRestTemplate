package auth

import (
	"github.com/gin-gonic/gin"
	"code/core/middleware"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.GET("/login", LoginHandler)
	rg.POST("/signup", middleware.ErrorWrapper(SignUpHandler))
	
}

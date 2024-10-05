package auth

import (
	"github.com/gin-gonic/gin"
	"code/core/middleware"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", middleware.ErrorWrapper(LoginHandler))
	rg.POST("/signup", middleware.ErrorWrapper(SignUpHandler))
}

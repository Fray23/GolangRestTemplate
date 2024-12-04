package routes

import (
	"code/api/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	usersGroup := r.Group("api/users")

	// ChangeMe
	// someApiGroup.Use(middleware.JwtMiddleware())

	auth.RegisterUserRoutes(usersGroup)
}

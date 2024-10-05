package routes

import (
	auth "code/api/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	auth.RegisterUserRoutes(r.Group("api/users"))
}

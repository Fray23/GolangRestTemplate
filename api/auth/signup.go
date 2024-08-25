package auth

import (
	dto "code/dto/api"
	service "code/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	var signUpData dto.SignUpDTO

	if err := c.BindJSON(&signUpData); err != nil {
		return
	}
	service.RegisterNewUser(&signUpData)

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully!",
		"test":    "another message",
	})
}

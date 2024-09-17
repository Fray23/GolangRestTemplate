package auth

import (
	"net/http"

	dto "code/dto/api"
	service "code/service"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	var signUpData dto.SignUpDTO

	if err := c.BindJSON(&signUpData); err != nil {
		return
	}
	err := service.RegisterNewUser(&signUpData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, dto.SignUpResponse{
		Status: 0,
		Message: "User created successfully!",
	})
}

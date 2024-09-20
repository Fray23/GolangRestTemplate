package auth

import (
	dto "code/dto/api"
	service "code/service"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) (*dto.SignUpResponse, error) {
	var signUpData dto.SignUpDTO

	if err := c.BindJSON(&signUpData); err != nil {
		return nil, err
	}
	err := service.RegisterNewUser(&signUpData)
	if err != nil {
		return nil, err
	}

	return &dto.SignUpResponse{
		Status: 0,
		Message: "User created successfully!",
	}, nil
}

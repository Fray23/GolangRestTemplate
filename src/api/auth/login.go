package auth

import (
	dto "code/core/dto/api"
	"github.com/gin-gonic/gin"
	"code/service"
)

func LoginHandler(c *gin.Context) (*dto.LoginResponse, error) {
	var AuthData dto.AuthDTO
	if err := c.BindJSON(&AuthData); err != nil {
		return nil, err
	}

	err := service.AuthUser(&AuthData)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Status: 0,
		Message: "auth",
	}, nil
}

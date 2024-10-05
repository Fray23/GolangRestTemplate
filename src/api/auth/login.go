package auth

import (
	dto "code/core/dto/api"
	"code/service"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) (*dto.LoginResponse, error) {
	var AuthData dto.AuthDTO
	if err := c.BindJSON(&AuthData); err != nil {
		return nil, err
	}

	token, err := service.JwtAuthUser(&AuthData)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Status:  0,
		Message: token,
	}, nil
}

package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	resut := gin.H{
		"message": "User created successfully!",
		"test":    "another message",
	}
	c.JSON(http.StatusOK, resut)
}

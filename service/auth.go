package service

import (
	"fmt"

	db "code/core/db"
	dto "code/dto/api"
	models "code/models"
)

func RegisterNewUser(signUpData *dto.SignUpDTO) {
	var user models.User
	if db.DB.First(&user, "login = ?", signUpData.Login).RecordNotFound() {
		fmt.Println("user with this login not exists")
	}
}

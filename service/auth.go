package service

import (
	db "code/core/db"
	dto "code/core/dto/api"
	repository "code/repository"
	auth_err "code/core/enum"
)

func RegisterNewUser(signUpData *dto.SignUpDTO) error {
	user_repository := repository.UserRepository{db.DB}
	operationResult := user_repository.GetUserByLogin(signUpData.Login)

	if operationResult.Error == nil {
		return auth_err.UserAlreadyRegistered
	} else {
		_ = user_repository.CreateUser(
			signUpData.Login,
			signUpData.Password,
		)
	}

	return nil
}

package service

import (
	db "code/core/db"
	dto "code/core/dto/api"
	repository "code/repository"
	auth_err "code/core/enum"
	"fmt"
	"code/core/security"
)

func RegisterNewUser(signUpData *dto.SignUpDTO) error {
	user_repository := repository.UserRepository{db.DB}
	operationResult := user_repository.GetUserByLogin(signUpData.Login)
	hash_password, err := security.HashPassword(signUpData.Password)

	if err != nil {
    	return fmt.Errorf("failed to generate hash for password: %w", err)
	}

	if operationResult.Error == nil {
		return auth_err.UserAlreadyRegistered
	} else {
		_ = user_repository.CreateUser(
			signUpData.Login,
			hash_password,
		)
	}

	return nil
}

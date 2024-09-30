package service

import (
	"fmt"

	db "code/core/db"
	dto "code/core/dto/api"
	auth_err "code/core/enum"
	"code/core/security"
	"code/models"
	repository "code/repository"
)

type jwt_token string

func RegisterNewUser(signUpData *dto.SignUpDTO) error {
	user_repository := repository.UserRepository{DB: db.DB}
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

func AuthUser(loginData *dto.AuthDTO) error {
	user_repository := repository.UserRepository{DB: db.DB}
	operationResult := user_repository.GetUserByLogin(loginData.Login)

	if operationResult.Error != nil {
		return auth_err.UserNotFound
	}

	user := operationResult.Result.(*db_models.User)

	if security.CheckPasswordHash(loginData.Password, user.Password) {
		return nil
	} else {
		return auth_err.InvalidPassword
	}
}

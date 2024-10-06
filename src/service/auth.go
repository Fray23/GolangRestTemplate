package service

import (
	"fmt"

	db "code/core/db"
	dto "code/core/dto/api"
	auth_err "code/core/enum"
	"code/core/security"
	db_models "code/models"
	repository "code/repository"
)

type jwt_token string

func RegisterNewUser(signUpData *dto.SignUpDTO) error {
	user_repository := repository.UserRepository{DB: db.DB}
	userAlreadyExists, err := user_repository.IsExistsByLogin(signUpData.Login)
	if err != nil {
		return fmt.Errorf("failed to check user is Exists: %w", err)
	}

	hash_password, err := security.HashPassword(signUpData.Password)
	if err != nil {
		return fmt.Errorf("failed to generate hash for password: %w", err)
	}

	if userAlreadyExists {
		return auth_err.UserAlreadyRegistered
	}

	_ = user_repository.CreateUser(
		signUpData.Login,
		hash_password,
	)

	return nil
}

func JwtAuthUser(loginData *dto.AuthDTO) (string, error) {
	user_repository := repository.UserRepository{DB: db.DB}
	operationResult := user_repository.GetUserByLogin(loginData.Login)

	if operationResult.Error != nil {
		return "", auth_err.UserNotFound
	}

	user := operationResult.Result.(*db_models.User)

	if !security.CheckPasswordHash(loginData.Password, user.Password) {
		return "", auth_err.InvalidPassword
	}

	token, err := security.GenerateJwtToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

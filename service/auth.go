package service

import (
	"fmt"

	db "code/core/db"
	dto "code/dto/api"
	db_models "code/models"
	repository "code/repository"
	auth_err "code/core/enum"
)

func RegisterNewUser(signUpData *dto.SignUpDTO) error {
	user_repository := repository.UserRepository{db.DB}
	operationResult := user_repository.GetUserByLogin(signUpData.Login)

	if operationResult.Error == nil {
		return &auth_err.ClientError{
			Message: "user already registered",
			Code: auth_err.USER_ALREADY_REGISTERED,
		}
	} else {
		user_result := user_repository.CreateUser(
			signUpData.Login,
			signUpData.Password,
		)

		user := user_result.Result.(*db_models.User)
		fmt.Println(user.Login)
	}

	return nil
}

package user_repo

import (
	dto "code/core/dto"
	db_models "code/models"

	"gorm.io/gorm"
)

type UserInterface interface {
	GetUserByID(id uint) (*db_models.User, error)
	CreateUser(user *db_models.User) error
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetUserByID(id uint) dto.UserRepositoryResult {
	var user db_models.User
	err := r.DB.First(&user, id).Take(&user).Error
	if err != nil {
		return dto.UserRepositoryResult{Error: err}
	}

	return dto.UserRepositoryResult{Result: &user}
}

func (r *UserRepository) IsExistsByLogin(login string) (bool, error) {
	user := &db_models.User{}
	err := r.DB.Select("id").Where("login = ?", login).Limit(1).Find(user).Error
	if err != nil {
		return false, err
	}

	if (*user == db_models.User{}) {
		return false, nil
	}
	return true, nil
}

func (r *UserRepository) GetUserByLogin(login string) dto.UserRepositoryResult {
	var user db_models.User
	err := r.DB.Where("login = ?", login).First(&user).Error
	if err != nil {
		return dto.UserRepositoryResult{Error: err}
	}
	return dto.UserRepositoryResult{Result: &user}
}

func (r *UserRepository) CreateUser(login string, password_hash string) dto.UserRepositoryResult {
	user := db_models.User{
		Login:    login,
		Password: password_hash,
	}
	result := r.DB.Create(&user)
	if result.Error != nil {
		return dto.UserRepositoryResult{Error: result.Error}
	}
	return dto.UserRepositoryResult{Result: &user}
}

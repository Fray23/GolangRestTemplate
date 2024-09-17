package user_repo

import (
	dto "code/dto"
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

// func NewUserRepository(db *gorm.DB) UserRepository {
// 	return &userRepository{db: db}
// }

func (r *UserRepository) GetUserByID(id uint) dto.RepositoryResult {
	var user db_models.User
	err := r.DB.First(&user, id).Take(&user).Error
	if err != nil {
		return dto.RepositoryResult{Error: err}
	}

	return dto.RepositoryResult{Result: &user}
}

func (r *UserRepository) GetUserByLogin(login string) dto.RepositoryResult {
	var user db_models.User
	err := r.DB.Where("login = ?", login).First(&user).Error
	if err != nil {
		return dto.RepositoryResult{Error: err}
	}
	return dto.RepositoryResult{Result: &user}
}

func (r *UserRepository) CreateUser(login string, password_hash string) dto.RepositoryResult {
	user := db_models.User{
		Login:    login,
		Password: password_hash,
	}
	result := r.DB.Create(&user)
	if result.Error != nil {
		return dto.RepositoryResult{Error: result.Error}
	}
	return dto.RepositoryResult{Result: &user}
}

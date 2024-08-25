package user_repo

import (
	db_models "code/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id uint) (*db_models.User, error)
	CreateUser(user *db_models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

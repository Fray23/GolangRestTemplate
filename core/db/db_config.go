package core

import (
	db_models "code/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := "user=postgres password=postgres dbname=gorm1 host=db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Migrate() error {
	err := DB.AutoMigrate(&db_models.User{})
	if err != nil {
		return err
	}
	return nil
}

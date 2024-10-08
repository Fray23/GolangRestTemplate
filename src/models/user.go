package db_models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Login     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

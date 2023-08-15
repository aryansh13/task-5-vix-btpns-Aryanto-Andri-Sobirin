package database

import (
	"go_restapi_gin/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Photo{})
	return err
}

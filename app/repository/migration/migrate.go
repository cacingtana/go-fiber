package migration

import (
	"go-fiber/app/models"

	"gorm.io/gorm"
)

func initMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Diary{}, &models.User{})
}

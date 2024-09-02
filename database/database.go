package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"josk/task-management-system/models"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	errDB := DB.AutoMigrate(&models.Task{})
	if errDB != nil {
		return
	}
}

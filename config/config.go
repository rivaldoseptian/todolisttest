package config

import (
	"log"
	"os"
	"todolist-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadEnv() {
	os.Setenv("PORT", "8080")
}

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DB_DSN")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DB.AutoMigrate(&models.User{}, &models.Checklist{}, &models.ChecklistItem{})
}

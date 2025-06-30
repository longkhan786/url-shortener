package db

import (
	"log"

	"github.com/longkhan786/url-shortener/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("shortener.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = DB.AutoMigrate(&models.URL{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}
}

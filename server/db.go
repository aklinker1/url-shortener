package server

import (
	"fmt"

	"github.com/aklinker1/url-shortener/server/models"
	"github.com/aklinker1/url-shortener/server/repos"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDB() {
	db, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to Postgres:", err)
	}
	DB = db
	err = DB.AutoMigrate(&models.URLEntry{})
	if err != nil {
		fmt.Println("Failed to migrate tables:", err)
	}

	// Init Repos
	repos.URLEntryRepo.Init(DB)
}

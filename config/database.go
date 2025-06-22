package config

import (
	"fmt"
	"log"
	"os"

	"hospital-system/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_URL") // DB_URL will be in .env
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&models.Patient{}) // Create table
	DB = db
	fmt.Println("Database connected successfully")
}

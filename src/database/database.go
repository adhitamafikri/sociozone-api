package database

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func Initialize() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
}

package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var connection *gorm.DB

// func Initialize() (*gorm.DB, error) {
func Initialize() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	return connectionString
}

func GetConnection() {
	connectionString := Initialize()
	fmt.Println(connectionString)
}

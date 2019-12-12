package main

import (
	"fmt"
	"os"

	database "github.com/adhitamafikri/sociozone-api/src/database/postgres"
	"github.com/adhitamafikri/sociozone-api/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var dbhost = os.Getenv("DB_NAME")
	fmt.Println("Welcome to Sociozone API", dbhost)

	rt := routes.RouteLoader{}

	database.GetConnection()
	// Initiate Gin/App
	router := gin.Default()

	// Load application routes
	rt.LoadRoutes(router)

	router.Run(":4001")
}

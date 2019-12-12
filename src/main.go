package main

import (
	"fmt"
	"os"

	"github.com/adhitamafikri/sociozone-api/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var dbhost = os.Getenv("DB_NAME")
	fmt.Println("Welcome to Sociozone API", dbhost)

	rt := routes.RouteLoader{}

	// Initiate Gin/App
	router := gin.Default()
	rt.LoadRoutes(router)

	// controllers.AuthControllerHandler(router)
	// controllers.UserControllerHandler(router)
	// controllers.PostsControllerHandler(router)

	router.Run(":4001")
}

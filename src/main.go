package main

import (
	"fmt"
	"os"

	"github.com/adhitamafikri/sociozone-api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var dbhost = os.Getenv("DB_NAME")
	fmt.Println("Welcome to Sociozone API", dbhost)

	// Initiate iris
	router := gin.Default()

	controllers.AuthControllerHandler(router)
	controllers.UserControllerHandler(router)
	controllers.PostsControllerHandler(router)

	router.Run(":4001")
}

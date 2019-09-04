package main

import (
	"fmt"

	"github.com/adhitamafikri/sociozone-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Sociozone API")

	// Initiate iris
	router := gin.Default()

	controllers.UserControllerHandler(router)
	controllers.AuthControllerHandler(router)

	router.Run(":4001")
}

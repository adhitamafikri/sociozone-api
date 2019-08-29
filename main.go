package main

import (
	"fmt"

	"github.com/adhitamafikri/sociozone-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Sociozone API")

	// Initiate iris
	app := gin.Default()
	routes.CreateRouteGroup(app)

	app.Run(":4001")
}

package main

import (
	"fmt"

	"github.com/adhitamafikri/sociozone-api/routes"
	"github.com/kataras/iris"
)

func main() {
	fmt.Println("Welcome to Sociozone API")

	// Initiate iris
	app := iris.Default()

	routes.CreateRouteGroup(app)

	// fmt.Println(app.GetRoutes())

	app.Run(iris.Addr(":4001"))
}

func getHome(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":  200,
		"message": "Welcome to home!",
	})
}

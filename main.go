package main

import (
	"fmt"

	"github.com/adhitamafikri/sociozone-api/constants"
	"github.com/kataras/iris"

	LoginController "github.com/adhitamafikri/sociozone-api/controllers/login"
	RegisterController "github.com/adhitamafikri/sociozone-api/controllers/register"
	UserController "github.com/adhitamafikri/sociozone-api/controllers/user"
)

func main() {
	fmt.Println("Welcome to Sociozone API")
	fmt.Println("base URL is ", constants.APIURL)

	// Initiate iris
	app := iris.Default()

	v1 := app.Party(constants.APIURL + "/v1")
	{
		v1.Get("/", getHome)

		v1.Get("/login", LoginController.Get)
		v1.Post("/login", LoginController.Post)

		v1.Get("/register", RegisterController.Get)
		v1.Post("/register", RegisterController.Post)

		v1.Get("/users", UserController.Get)
	}

	app.Run(iris.Addr(":4001"))
}

func getHome(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":  200,
		"message": "Welcome to home!",
	})
}

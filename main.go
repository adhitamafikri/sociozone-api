package main

import (
	"fmt"

	"github.com/adhitamafikri/sociozone-api/constants"
	"github.com/kataras/iris"

	LoginController "github.com/adhitamafikri/sociozone-api/controllers/login"
	RegisterController "github.com/adhitamafikri/sociozone-api/controllers/register"
	UserController "github.com/adhitamafikri/sociozone-api/controllers/user"

	DBHelper "github.com/adhitamafikri/sociozone-api/helpers/db"
)

func main() {
	fmt.Println("Welcome to Sociozone API")
	fmt.Println("base URL is ", constants.APIURL)

	// Establish DB connection
	DBHelper.ConnectDB()

	// Initiate iris
	app := iris.Default()

	v1 := app.Party(constants.APIURL + "/v1")
	{
		v1.Get("/", getHome)
		v1.Get("/login", LoginController.Index)
		v1.Post("/login", LoginController.Post)
		v1.Get("/register", RegisterController.Index)
		v1.Post("/register", RegisterController.Post)
		v1.Get("/users", UserController.Index)
	}

	app.Run(iris.Addr(":4001"))
}

func getHome(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":  200,
		"message": "Welcome to home!",
	})
}

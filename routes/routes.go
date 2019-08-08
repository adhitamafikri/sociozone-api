package routes

import (
	"fmt"

	"github.com/adhitamafikri/sociozone-api/constants"
	"github.com/kataras/iris"

	LoginController "github.com/adhitamafikri/sociozone-api/controllers/login"
	PostsController "github.com/adhitamafikri/sociozone-api/controllers/posts"
	RegisterController "github.com/adhitamafikri/sociozone-api/controllers/register"
	UserController "github.com/adhitamafikri/sociozone-api/controllers/user"
)

// CreateRouteGroup returns all available v1 for this API
func CreateRouteGroup(app *iris.Application) {
	fmt.Println("from v1.go, Base URL is ", constants.APIURL)

	v1 := app.Party(constants.APIURL + "/v1")
	{
		v1.Get("/", getHome)

		v1.Get("/login", LoginController.Get)
		v1.Post("/login", LoginController.Post)

		v1.Get("/register", RegisterController.Get)
		v1.Post("/register", RegisterController.Post)

		v1.Get("/users", UserController.Get)

		v1.Get("/posts", PostsController.Get)
		v1.Post("/posts", PostsController.Post)
	}
}

func getHome(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":  200,
		"message": "Welcome to home!",
	})
}

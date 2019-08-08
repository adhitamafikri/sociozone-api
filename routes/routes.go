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

// RouteGroup returns all available routes for this API
func RouteGroup(app *iris.Application) {
	fmt.Println("from routes.go, Base URL is ", constants.APIURL)

	routes := app.Party(constants.APIURL + "/v1")
	{
		routes.Get("/", getHome)

		routes.Get("/login", LoginController.Get)
		routes.Post("/login", LoginController.Post)

		routes.Get("/register", RegisterController.Get)
		routes.Post("/register", RegisterController.Post)

		routes.Get("/users", UserController.Get)

		routes.Get("/posts", PostsController.Get)
	}
}

func getHome(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":  200,
		"message": "Welcome to home!",
	})
}

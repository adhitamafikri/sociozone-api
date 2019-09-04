package routes

import (
	"fmt"
	"net/http"

	"github.com/adhitamafikri/sociozone-api/constants"
	"github.com/gin-gonic/gin"

	// LoginController "github.com/adhitamafikri/sociozone-api/controllers/login"
	// PostsController "github.com/adhitamafikri/sociozone-api/controllers/posts"
	// RegisterController "github.com/adhitamafikri/sociozone-api/controllers/register"
	UserController "github.com/adhitamafikri/sociozone-api/controllers/user"
)

// CreateRouteGroup returns all available v1 for this API
func CreateRouteGroup(app *gin.Engine) {
	fmt.Println("from v1.go, Base URL is ", constants.APIURL)

	v1 := app.Group(constants.APIURL + "/v1")
	{
		// v1.GET("/", getHome)
		//
		// v1.GET("/login", LoginController.Get)
		// v1.POST("/login", LoginController.Post)
		//
		// v1.GET("/register", RegisterController.Get)
		// v1.POST("/register", RegisterController.Post)

		v1.GET("/users", UserController.Get)

		// v1.GET("/posts", PostsController.Get)
		// v1.POST("/posts", PostsController.Post)
	}
}

func getHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to home!",
	})
}

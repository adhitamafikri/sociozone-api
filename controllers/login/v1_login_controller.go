package controllers

import (
	LoginService "github.com/adhitamafikri/sociozone-api/services/login"
	"github.com/kataras/iris"
)

// Get will open login screen
func Get(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":  200,
		"message": "Opening login screen",
	})
}

// Post will post login data
func Post(ctx iris.Context) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	LoginService.Login(&username, &password)

	ctx.JSON(iris.Map{
		"status":  200,
		"message": "Posting Login Data",
	})
}

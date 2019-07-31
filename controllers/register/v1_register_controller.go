package controllers

import (
	RegisterService "github.com/adhitamafikri/sociozone-api/services/register"
	"github.com/kataras/iris"
)

// Index will open registration form
func Index(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":  200,
		"message": "opening registration form",
	})
}

// Post will post registration data to DB
func Post(ctx iris.Context) {
	name := ctx.FormValue("name")
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	RegisterService.RegisterUserData(&name, &username, &password)

	ctx.JSON(iris.Map{
		"status":   200,
		"message":  "Posting Registration form",
		"name":     name,
		"username": username,
		"password": password,
	})
}

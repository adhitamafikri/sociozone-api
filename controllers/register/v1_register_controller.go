package controllers

import (
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

	ctx.JSON(iris.Map{
		"status":   200,
		"message":  "Posting Registration form",
		"name":     name,
		"username": username,
		"password": password,
	})
}

package controllers

import (
	"github.com/kataras/iris"
)

// Index will open login screen
func Index(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":  200,
		"message": "Opening login screen",
	})
}

// Post will post login data
func Post(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":  200,
		"message": "Posting Login Data",
	})
}

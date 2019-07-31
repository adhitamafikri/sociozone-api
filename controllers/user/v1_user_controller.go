package controllers

import (
	"fmt"

	UserService "github.com/adhitamafikri/sociozone-api/services/user"
	"github.com/kataras/iris"
)

// Index retrieves all users
func Index(ctx iris.Context) {
	fmt.Println("Getting all users from controlerr")
	UserService.Index()

	ctx.JSON(iris.Map{
		"status":  200,
		"message": "Done retrieving users",
	})
}

package controllers

import (
	"fmt"

	PostsService "github.com/adhitamafikri/sociozone-api/services/posts"
	"github.com/kataras/iris"
)

// Get retrieves all posts
func Get(ctx iris.Context) {
	fmt.Println("Getting all posts")

	PostsService.RetrievePosts()

	ctx.JSON(iris.Map{
		"status":  200,
		"message": "Getting all posts",
	})
}
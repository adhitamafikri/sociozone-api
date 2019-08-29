package controllers

import (
	"fmt"
	"net/http"

	PostsService "github.com/adhitamafikri/sociozone-api/services/posts"
	"github.com/gin-gonic/gin"
)

// Get retrieves all posts
func Get(ctx *gin.Context) {
	fmt.Println("Getting all posts")

	PostsService.RetrievePosts()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Getting all posts",
	})
}

// Post stores user's post into DB
func Post(ctx *gin.Context) {
	fmt.Println("Posting user post")

	PostsService.UploadPost()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Posting user post",
	})
}

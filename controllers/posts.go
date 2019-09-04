package controllers

import (
	"fmt"
	"net/http"

	"github.com/adhitamafikri/sociozone-api/services"
	"github.com/gin-gonic/gin"
)

// Get retrieves all posts
func GetPosts(ctx *gin.Context) {
	fmt.Println("Getting all posts")

	services.RetrievePosts()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Getting all posts",
	})
}

// Post stores user's post into DB
func PostPosts(ctx *gin.Context) {
	fmt.Println("Posting user post")

	services.UploadPost()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Posting user post",
	})
}

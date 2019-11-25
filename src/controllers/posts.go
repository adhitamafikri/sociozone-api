package controllers

import (
	"fmt"
	"net/http"

	"github.com/adhitamafikri/sociozone-api/constants/config"
	"github.com/adhitamafikri/sociozone-api/services"
	"github.com/gin-gonic/gin"
)

type PostsController struct {
	service services.PostsService
}

func PostsControllerHandler(router *gin.Engine) {
	handler := &PostsController{service: services.PostsServiceHandler()}

	group := router.Group(config.APIURL + "/v1/posts")
	group.GET("/", handler.GetAllPosts)
}

// Get retrieves all posts
func (handler *PostsController) GetAllPosts(ctx *gin.Context) {
	fmt.Println("Getting all posts")

	handler.service.GetAllPosts()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Getting all posts",
	})
}

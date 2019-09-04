package controllers

import (
	"fmt"
	"net/http"

	LoginService "github.com/adhitamafikri/sociozone-api/services/login"
	"github.com/gin-gonic/gin"
)

// Get will open login screen
func Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Opening login screen",
	})
}

// Post will post login data
func Post(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username, password)

	LoginService.Login(username, password)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Posting Login Data",
	})
}

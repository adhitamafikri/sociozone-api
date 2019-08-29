package controllers

import (
	"net/http"

	RegisterService "github.com/adhitamafikri/sociozone-api/services/register"
	"github.com/gin-gonic/gin"
)

// Get will open registration form
func Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "opening registration form",
	})
}

// Post will post registration data to DB
func Post(ctx *gin.Context) {
	name := ctx.PostForm("name")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	RegisterService.RegisterUserData(name, username, password)

	ctx.JSON(http.StatusOK, gin.H{
		"status":   200,
		"message":  "Posting Registration form",
		"name":     name,
		"username": username,
		"password": password,
	})
}

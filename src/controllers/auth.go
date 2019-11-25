package controllers

import (
	"fmt"
	"net/http"

	"github.com/adhitamafikri/sociozone-api/src/constants/config"
	"github.com/adhitamafikri/sociozone-api/src/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service services.AuthService
}

func AuthControllerHandler(router *gin.Engine) {
	handler := &AuthController{service: services.AuthServiceHandler()}

	group := router.Group(config.APIURL + "/v1/auth")
	group.GET("/login", handler.GetLogin)
	group.POST("/login", handler.PostLogin)
	group.GET("/register", handler.GetRegister)
	group.POST("/register", handler.PostRegister)
}

func (handler *AuthController) GetLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Getting Login Page",
	})
}

func (handler *AuthController) PostLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username, password)

	handler.service.PostLogin(username, password)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Posting Login Data",
	})
}

func (handler *AuthController) GetRegister(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Getting Register Page",
	})
}

func (handler *AuthController) PostRegister(ctx *gin.Context) {
	name := ctx.PostForm("name")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	handler.service.PostRegister(name, username, password)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Posting Registration Data",
	})
}

package controllers

import (
	"fmt"
	"net/http"

	"github.com/adhitamafikri/sociozone-api/constants/config"
	"github.com/adhitamafikri/sociozone-api/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func UserControllerHandler(router *gin.Engine) {
	handler := &UserController{service: services.UserServiceHandler()}

	group := router.Group(config.APIURL + "/v1/users")
	group.GET("/", handler.GetAllUsers)
}

// Get retrieves all users
func (handler *UserController) GetAllUsers(ctx *gin.Context) {
	fmt.Println("Getting all users from controlerr")
	handler.service.GetAllUsers()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Done retrieving users",
	})
}

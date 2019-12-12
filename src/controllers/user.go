package controllers

import (
	"github.com/adhitamafikri/sociozone-api/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserController is controller for user :)
type UserController struct {
	UserService services.UserService
}

type iUserController interface {
	GetAllUsers()
}

func (ctrl *UserController) GetAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Getting all users",
	})
}

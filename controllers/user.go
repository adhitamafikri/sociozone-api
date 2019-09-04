package controllers

import (
	"fmt"
	"net/http"

	"github.com/adhitamafikri/sociozone-api/services/users"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	service services.UserService
}

// Get retrieves all users
func (ctrl *UsersController) Get(ctx *gin.Context) {
	fmt.Println("Getting all users from controlerr")
	ctrl.service.Index()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Done retrieving users",
	})
}

package controllers

import (
	"fmt"
	"net/http"

	UserService "github.com/adhitamafikri/sociozone-api/services/user"
	"github.com/gin-gonic/gin"
)

// Get retrieves all users
func Get(ctx *gin.Context) {
	fmt.Println("Getting all users from controlerr")
	UserService.Index()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Done retrieving users",
	})
}

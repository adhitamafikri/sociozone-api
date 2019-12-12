package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/adhitamafikri/sociozone-api/src/controllers"
	"github.com/adhitamafikri/sociozone-api/src/services"
)

// RouteLoader is a struct to encapsulate LoadRoutes() method
type RouteLoader struct{}

// LoadRoutes is a method to generate routes of the app
func (rLoader *RouteLoader) LoadRoutes(router *gin.Engine) {
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Skrrrt",
		})
	})

	userHandler := &controllers.UserController{
		UserService: services.UserServiceHandler(),
	}

	// Route generations
	rLoader.routeUsers(router, userHandler)
	// Route generations: END
}

func (rLoader *RouteLoader) routeUsers(router *gin.Engine, handler *controllers.UserController) {
	// User routes group
	groupUser := router.Group("/v1/users")
	groupUser.GET("/", handler.GetAllUsers)
	// User routes group: END
}

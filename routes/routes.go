package routes

import (
	"github.com/fabianoflorentino/gotostudy/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	user := router.Group("/users")
	{
		user.GET("", controllers.GetUsers)
		user.GET("/:id", controllers.GetUserByID)
		user.POST("", controllers.CreateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	health := router.Group("/health")
	{
		health.GET("", controllers.HealthCheck)
	}
}

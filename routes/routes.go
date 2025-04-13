package routes

import (
	"github.com/fabianoflorentino/gotostudy/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	r := router.Group("/users")
	{
		r.GET("", controllers.GetUsers)
	}

	h := router.Group("/health")
	{
		h.GET("", controllers.HealthCheck)
	}
}

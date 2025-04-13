// Package routes defines the HTTP routes and their corresponding handlers
// for the application. It organizes the endpoints into groups and associates
// them with the appropriate controller functions to handle incoming requests.
package routes

import (
	"github.com/fabianoflorentino/gotostudy/controllers"
	"github.com/gin-gonic/gin"
)

// InitializeRoutes sets up the API routes for the application.
// It defines route groups and their respective endpoints, associating them
// with the appropriate controller functions.
//
// Parameters:
//   - router (*gin.Engine): The Gin engine instance used to define the routes.
//
// Routes:
//   - /users:
//   - GET    "": Retrieves a list of users.
//   - GET    "/:id": Retrieves a user by their ID.
//   - POST   "": Creates a new user.
//   - PUT    "/:id": Updates an existing user by their ID.
//   - PATCH  "/:id": Partially updates fields of an existing user by their ID.
//   - DELETE "/:id": Deletes a user by their ID.
//   - /health:
//   - GET    "": Performs a health check of the application.
func InitializeRoutes(router *gin.Engine) {
	user := router.Group("/users")
	{
		user.GET("", controllers.GetUsers)
		user.GET("/:id", controllers.GetUserByID)
		user.POST("", controllers.CreateUser)
		user.PUT("/:id", controllers.UpdateUser)
		user.PATCH("/:id", controllers.UpdateUserFields)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	health := router.Group("/health")
	{
		health.GET("", controllers.HealthCheck)
	}
}

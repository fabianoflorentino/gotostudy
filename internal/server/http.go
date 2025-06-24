package server

import (
	"log"
	"os"

	"github.com/fabianoflorentino/gotostudy/adapters/inbound/http/controllers"
	"github.com/fabianoflorentino/gotostudy/internal/app"
	"github.com/gin-gonic/gin"
)

// StartHTTPServer initializes a new Gin HTTP server with the specified configuration.
// It sets the server to run in release mode, configures trusted proxies,
// and sets up the router with the provided controller.
func StartHTTPServer(container *app.AppContainer) {
	r := gin.Default()

	setTrustedProxies(r)

	registerUserRoutes(r, container)
	registerTaskRoutes(r, container)
	registerHealthRoutes(r)

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Printf("Failed to start HTTP server: %v", err)
	}
}

// RegisterUserRoutes sets up the user-related routes for the Gin HTTP server.
// It registers the routes for creating a user, getting all users, and getting a user by ID.
func registerUserRoutes(r *gin.Engine, container *app.AppContainer) {
	userController := controllers.NewUserController(container.UserService)

	r.POST("/users", userController.CreateUser)
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUserByID)
	r.PUT("/users/:id", userController.UpdateUser)
	r.PATCH("/users/:id", userController.UpdateUserFields)
	r.DELETE("/users/:id", userController.DeleteUser)
}

// RegisterTaskRoutes sets up the task-related routes for the Gin HTTP server.
func registerTaskRoutes(r *gin.Engine, container *app.AppContainer) {
	taskController := controllers.NewTaskController(container.TaskService)

	r.POST("/users/:id/tasks", taskController.CreateTask)
	r.GET("/users/:id/tasks", taskController.FindUserTasks)
	r.GET("/users/:id/tasks/:task_id", taskController.FindTaskByID)
	// r.PUT("/tasks/:id", taskController.UpdateTask)
	// r.PATCH("/tasks/:id", taskController.UpdateTaskFields)
	// r.DELETE("/tasks/:id", taskController.DeleteTask)
}

// RegisterHealthRoutes sets up the health check route for the Gin HTTP server.
// It registers a route to check the health of the application.
func registerHealthRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}

// SetTrustedProxies configures the trusted proxies for the Gin HTTP server.
// It sets the trusted proxies to allow the server to correctly handle forwarded headers.
func setTrustedProxies(r *gin.Engine) {
	trustedProxies := []string{"127.0.0.1", "::1", "192.168.0.0/16", "172.16.0.0/8"}

	if err := r.SetTrustedProxies(trustedProxies); err != nil {
		log.Printf("Failed to set trusted proxies: %v", err)
	}
}

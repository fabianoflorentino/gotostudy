// File: main.go
// Description: This is the main entry point for the GoToStudy application.
// It initializes the application by loading environment variables, setting up the database,
// and configuring the HTTP server with routes.
package main

import (
	"log"

	"github.com/fabianoflorentino/gotostudy/config"
	"github.com/fabianoflorentino/gotostudy/database"
	"github.com/fabianoflorentino/gotostudy/internal/http_config"
	"github.com/fabianoflorentino/gotostudy/routes"
	"github.com/gin-gonic/gin"
)

// init initializes the application by loading environment variables and initializing the database.
// It is called before the main function.
func init() {
	config.LoadEnv()
	database.InitDB()
}

// main is the entry point of the application.
// It sets up the Gin router, configures trusted proxies, and initializes routes.
// Finally, it starts the HTTP server.
func main() {
	r := gin.Default()

	http_config.SetTrustedProxies(r)
	routes.InitializeRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// File: main.go
// Description: This is the main entry point for the GoToStudy application.
// It initializes the application by loading environment variables, setting up the database,
// and configuring the HTTP server with routes.
package main

import (
	"log"

	"github.com/fabianoflorentino/gotostudy/internal/app"
	"github.com/fabianoflorentino/gotostudy/internal/server"
	"github.com/joho/godotenv"
)

// init initializes the application by loading environment variables and initializing the database.
// It is called before the main function.
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// main is the entry point of the application.
// It sets up the Gin router, configures trusted proxies, and initializes routes.
// Finally, it starts the HTTP server.
func main() {
	container := app.NewAppContainer()
	server.StartHTTPServer(container)
}

package server

import (
	"log"

	"github.com/fabianoflorentino/gotostudy/internal/app"
	"github.com/gin-gonic/gin"
)

// StartHTTPServer initializes a new Gin HTTP server with the specified configuration.
// It sets the server to run in release mode, configures trusted proxies,
// and sets up the router with the provided controller.
func StartHTTPServer(container *app.AppContainer) {
	r := gin.Default()

	setTrustedProxies(r)
	container.Controller.RegisterRoutes(r)

	log.Println("Starting HTTP server on port 8080")
	log.Fatal(r.Run(":8080"))
}

// SetTrustedProxies configures the trusted proxies for the Gin HTTP server.
// It sets the trusted proxies to allow the server to correctly handle forwarded headers.
func setTrustedProxies(r *gin.Engine) {
	trustedProxies := []string{"127.0.0.1", "::1", "192.168.0.0/16", "172.16.0.0/8"}

	if err := r.SetTrustedProxies(trustedProxies); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
}

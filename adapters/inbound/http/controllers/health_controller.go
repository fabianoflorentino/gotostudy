// Package controllers contains the HTTP handler functions and controllers
// responsible for processing incoming requests and returning responses.
// It serves as the inbound adapter in the application architecture,
// handling HTTP-specific logic and delegating business logic to the appropriate services.
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController is a struct that serves as a controller for handling
// health check-related HTTP requests. It is typically used to provide
// endpoints that allow clients to verify the application's availability
// and operational status.
type HealthController struct{}

// NewHealthController creates and returns a new instance of HealthController.
// This function initializes the HealthController struct and is typically used
// to set up the controller for handling health-related HTTP requests.
func NewHealthController() *HealthController {
	return &HealthController{}
}

// HealthCheck is a handler method for checking the health status of the application.
// It responds with an HTTP 200 status code and a JSON message indicating that the service is operational.
func (h *HealthController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// File: health_controller.go
// Description: This file contains the HealthCheck controller function.
// It handles the health check endpoint of the application.
package controllers

import (
	"net/http"

	"github.com/fabianoflorentino/gotostudy/services"
	"github.com/gin-gonic/gin"
)

// HealthCheck handles the health check endpoint.
// It retrieves the health status of the application and returns it as a JSON response.
func HealthCheck(c *gin.Context) {
	h, err := services.GetHealth()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, h)
}

package controllers

import (
	"net/http"

	"github.com/fabianoflorentino/gotostudy/services"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	h, err := services.GetHealth()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, h)
}

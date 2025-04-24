package handlers

import (
	"net/http"

	"github.com/fabianoflorentino/gotostudy/adapters/inbound/http/helpers"
	"github.com/fabianoflorentino/gotostudy/core/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ShouldBindJSON(c *gin.Context, input any) error {
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func HasValidUpdateUserFields(service *services.UserService, c *gin.Context, uid uuid.UUID) map[string]any {
	if helpers.UserExists(service, uid, c) == nil {
		return nil
	}

	updates, err := helpers.ParseUpdateFields(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	if !helpers.HasValidUpdates(updates, c) {
		return nil
	}

	return updates
}

// Package helpers provides utility functions to assist with common operations
// related to HTTP request handling and user management in the application.
// These helper functions include parsing UUIDs, validating and extracting
// update fields from JSON payloads, and checking the existence of users.
// The package is designed to streamline and centralize reusable logic for
// better maintainability and consistency across the codebase.
package helpers

import (
	"fmt"
	"net/http"

	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ParseUUID takes a string representation of a UUID and attempts to parse it into a uuid.UUID object.
// If the provided string is not a valid UUID, it returns an error indicating the issue.
// On success, it returns the parsed uuid.UUID and a nil error.
func ParseUUID(id string) (uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID: %s", err)
	}

	return uid, nil
}

// ParseUpdateFields parses and validates the JSON payload from the HTTP request context
// to extract fields for updating a user. It ensures that only valid fields are included
// in the update map. If the JSON payload contains invalid fields or cannot be bound,
// an error is returned.
//
// Parameters:
// - c: The Gin HTTP request context containing the JSON payload.
//
// Returns:
// - A map of valid fields to be updated and their corresponding values.
// - An error if the JSON payload is invalid or contains unsupported fields.
func ParseUpdateFields(c *gin.Context) (map[string]any, error) {
	var updates map[string]any

	if err := c.ShouldBindJSON(&updates); err != nil {
		return nil, err
	}

	validFields := map[string]bool{
		"username": true,
		"email":    true,
	}

	for field := range updates {
		if !validFields[field] {
			return nil, fmt.Errorf("invalid field: %s", field)
		}
	}

	return updates, nil
}

// HasValidUpdates checks if the provided updates map contains any valid fields to update.
// If the map is empty, it responds with a 400 Bad Request status and an error message
// indicating that there are no valid fields to update. Returns true if the updates map
// is not empty, otherwise returns false.
func HasValidUpdates(updates map[string]any, c *gin.Context) bool {
	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid fields to update"})
		return false
	}

	return true
}

// userExists checks if a user with the given UUID exists in the system.
// It retrieves the user by calling the service layer's GetUserByID method.
// If the user is not found or an error occurs during retrieval, it responds
// with an HTTP 404 status and a JSON error message. If the user exists,
// it returns the user object; otherwise, it returns nil.
func UserExists(service *services.UserService, uid uuid.UUID, c *gin.Context) *domain.User {
	user, err := service.GetUserByID(c, uid)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return nil
	}

	return user
}

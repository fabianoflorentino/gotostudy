// Package handlers provides HTTP handler functions for managing user-related
// operations in the application. These handlers interact with the core services
// and helpers to process incoming HTTP requests, validate input data, and
// return appropriate responses. The package is designed to work with the Gin
// web framework and includes utility functions for JSON binding, user existence
// checks, and validation of update fields.
package handlers

import (
	"github.com/gin-gonic/gin"
)

// ShouldBindJSON is a helper function that attempts to bind the JSON payload
// from the HTTP request body to the provided input structure. It uses the
// Gin framework's ShouldBindJSON method for binding. If the binding fails,
// it responds with a 400 Bad Request status and includes the error message
// in the response body. The function returns the error if binding fails,
// or nil if the binding is successful.
func ShouldBindJSON(c *gin.Context, input any) error {
	if err := c.ShouldBindJSON(input); err != nil {
		return err
	}

	return nil
}

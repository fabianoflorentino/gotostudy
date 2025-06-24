// Package helpers provides utility functions to assist with common operations
// related to HTTP request handling and user management in the application.
// These helper functions include parsing UUIDs, validating and extracting
// update fields from JSON payloads, and checking the existence of users.
// The package is designed to streamline and centralize reusable logic for
// better maintainability and consistency across the codebase.
package helpers

import (
	"fmt"

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

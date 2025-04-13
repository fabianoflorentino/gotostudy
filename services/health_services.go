// Package services provides the business logic and service layer for the application.
// It acts as an intermediary between the repositories and the controllers, ensuring
// that the application's core functionality is implemented and maintained.
package services

import "github.com/fabianoflorentino/gotostudy/repositories"

// GetHealth retrieves the health status of the application by calling the
// corresponding repository function. It returns a map containing health
// information and an error, if any occurs during the process.
func GetHealth() (map[string]string, error) {
	return repositories.GetHealth(), nil
}

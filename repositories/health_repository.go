// Package repositories provides the data access layer for the application.
// It contains functions and methods to interact with and retrieve data from
// various sources, such as databases or external APIs.
package repositories

// GetHealth returns a map containing the health status of the application.
// The map includes a single key-value pair where the key is "status" and the
// value is a message indicating that the application is alive and running.
func GetHealth() map[string]string {
	return map[string]string{
		"status": "I'm alive and running!",
	}
}

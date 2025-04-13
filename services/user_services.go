// Package services provides the business logic and service layer for the application.
// It acts as an intermediary between the repositories and the controllers,
// handling operations related to users such as retrieval, creation, updating, and deletion.
package services

import (
	"github.com/fabianoflorentino/gotostudy/models"
	"github.com/fabianoflorentino/gotostudy/repositories"
	"github.com/google/uuid"
)

// GetAllUsers retrieves all user records from the data repository.
// It returns a slice of models.User and an error if any issue occurs
// during the retrieval process.
func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

// GetUserByID retrieves a user from the repository based on the provided UUID.
//
// Parameters:
//   - id: The UUID of the user to be retrieved.
//
// Returns:
//   - models.User: The user object corresponding to the provided UUID.
//   - error: An error if the user could not be retrieved or does not exist.
func GetUserByID(id uuid.UUID) (models.User, error) {
	return repositories.GetUserByID(id)
}

// CreateUser creates a new user in the system by delegating the operation
// to the repositories layer. It takes a models.User object as input and
// returns the created user along with any error encountered during the process.
//
// Parameters:
//   - user: A models.User object containing the details of the user to be created.
//
// Returns:
//   - models.User: The newly created user object.
//   - error: An error object if the operation fails, otherwise nil.
func CreateUser(user models.User) (models.User, error) {
	return repositories.CreateUser(user)
}

// UpdateUser updates an existing user in the repository with the provided user data.
//
// Parameters:
//   - id: The UUID of the user to be updated.
//   - user: A models.User object containing the updated user information.
//
// Returns:
//   - models.User: The updated user object after the changes have been saved.
//   - error: An error object if the update operation fails, otherwise nil.
func UpdateUser(id uuid.UUID, user models.User) (models.User, error) {
	return repositories.UpdateUser(id, user)
}

// UpdateUserFields updates specific fields of a user identified by the given UUID.
// The fields to be updated are provided as a map where the keys are the field names
// and the values are the new values for those fields.
//
// Parameters:
//   - id: The UUID of the user whose fields are to be updated.
//   - fields: A map containing the field names as keys and their corresponding new values.
//
// Returns:
//   - models.User: The updated user object.
//   - error: An error if the update operation fails.
func UpdateUserFields(id uuid.UUID, fields map[string]any) (models.User, error) {
	return repositories.UpdateUserFields(id, fields)
}

// DeleteUser deletes a user from the repository based on the provided UUID.
//
// Parameters:
//   - id: The UUID of the user to be deleted.
//
// Returns:
//   - error: An error if the deletion fails, or nil if the operation is successful.
func DeleteUser(id uuid.UUID) error {
	return repositories.DeleteUser(id)
}

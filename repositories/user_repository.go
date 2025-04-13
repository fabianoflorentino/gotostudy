// Package repositories provides functions to interact with the database for managing user data.
// It includes operations such as retrieving all users, fetching a user by ID, creating a new user,
// updating user details, updating specific fields of a user, and deleting a user.
package repositories

import (
	"github.com/fabianoflorentino/gotostudy/database"
	"github.com/fabianoflorentino/gotostudy/models"
	"github.com/google/uuid"
)

var (
	users []models.User
	user  models.User
)

// GetAllUsers retrieves all user records from the database.
// It returns a slice of models.User and an error if any occurs during the database query.
func GetAllUsers() ([]models.User, error) {
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID retrieves a user from the database by their unique ID.
// It takes a UUID as input and returns the corresponding user model
// along with an error if the operation fails.
//
// Parameters:
//   - id: The UUID of the user to retrieve.
//
// Returns:
//   - models.User: The user model corresponding to the given ID.
//   - error: An error if the user cannot be found or if there is an issue
//     with the database operation.
func GetUserByID(id uuid.UUID) (models.User, error) {
	if err := database.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

// CreateUser creates a new user record in the database.
// It takes a models.User object as input and returns the created user object
// along with an error, if any occurs during the database operation.
//
// Parameters:
//   - user: The user object to be created in the database.
//
// Returns:
//   - models.User: The created user object with updated fields (e.g., ID).
//   - error: An error object if the operation fails, otherwise nil.
func CreateUser(user models.User) (models.User, error) {
	if err := database.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

// UpdateUser updates an existing user in the database with the provided user data.
// It takes a UUID representing the user's ID and a models.User object containing
// the updated user information. The function returns the updated user object and
// an error if the update operation fails.
//
// Parameters:
//   - id: The UUID of the user to be updated.
//   - user: A models.User object containing the updated user data.
//
// Returns:
//   - models.User: The updated user object.
//   - error: An error if the update operation fails, or nil if successful.
func UpdateUser(id uuid.UUID, user models.User) (models.User, error) {
	if err := database.DB.Save(&user).Where("id = ?", id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

// UpdateUserFields updates specific fields of a user in the database based on the provided user ID.
//
// Parameters:
//   - id: The UUID of the user whose fields are to be updated.
//   - fields: A map containing the fields to be updated and their new values.
//
// Returns:
//   - models.User: The updated user object.
//   - error: An error object if the operation fails, otherwise nil.
//
// The function performs the following steps:
//  1. Updates the specified fields of the user in the database.
//  2. Retrieves the updated user record from the database.
//  3. Returns the updated user object or an error if any operation fails.
func UpdateUserFields(id uuid.UUID, fields map[string]any) (models.User, error) {
	if err := database.DB.Model(&user).Where("id = ?", id).Updates(fields).Error; err != nil {
		return models.User{}, err
	}

	if err := database.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

// DeleteUser deletes a user from the database based on the provided UUID.
// It first retrieves the user record using the given ID and then deletes it.
// If any error occurs during the retrieval or deletion process, it returns the error.
//
// Parameters:
//   - id: The UUID of the user to be deleted.
//
// Returns:
//   - error: An error object if the operation fails, or nil if the operation is successful.
func DeleteUser(id uuid.UUID) error {
	if err := database.DB.First(&user, id).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

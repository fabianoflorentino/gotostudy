// Package services provides the implementation of business logic and
// application services for the GoToStudy application. It acts as an
// intermediary layer between the domain models and the repository layer,
// ensuring that the application's core functionalities are executed
// according to the business rules and requirements.
//
// The UserService struct in this package is responsible for managing
// user-related operations, such as registering new users, retrieving
// user information, updating user details, and deleting users. It
// interacts with the UserRepository interface to persist and retrieve
// user data from the underlying storage.
package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/fabianoflorentino/gotostudy/core"
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/ports"
	"github.com/google/uuid"
)

// UserService is a service layer struct that provides methods to manage user-related operations.
// It depends on a UserRepository interface (defined in the ports package) to interact with the underlying data storage.
type UserService struct {
	repo ports.UserRepository
}

// NewUserService creates and returns a new instance of UserService.
// It takes a UserRepository as a parameter, which is used to interact
// with the underlying data storage for user-related operations.
func NewUserService(r ports.UserRepository) *UserService {
	return &UserService{repo: r}
}

// RegisterUser creates a new user with the provided name and email, assigns a unique ID,
// and initializes an empty list of tasks for the user. It then saves the user to the repository.
// If the save operation fails, it logs the error and returns it. On success, it returns the created user.
func (s *UserService) RegisterUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	existingUser, err := s.repo.FindByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return nil, core.ErrUserAlreadyExists
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Tasks = []domain.Task{}

	if err := s.repo.Save(ctx, user); err != nil {
		log.Printf("Error saving user: %v", err)
		return nil, err
	}

	return user, nil
}

// GetAllUsers retrieves all users from the repository.
// It returns a slice of User objects and an error if any occurs during the retrieval process.
// If an error is encountered, it logs the error and returns nil along with the error.
func (s *UserService) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID retrieves a user from the repository based on the provided UUID.
// It returns a pointer to the User domain model if found, or an error if the user
// cannot be fetched or does not exist.
//
// Parameters:
//   - id: The UUID of the user to be retrieved.
//
// Returns:
//   - *domain.User: A pointer to the User object if found.
//   - error: An error object if there is an issue during retrieval.
func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if errors.Is(err, core.ErrUserNotFound) {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates an existing user in the repository with the provided user data.
// It takes a UUID representing the user's ID and a pointer to a domain.User object containing
// the updated user information. If the update operation fails, it logs the error and returns it.
// Otherwise, it returns nil to indicate success.
func (s *UserService) UpdateUser(ctx context.Context, id uuid.UUID, user *domain.User) error {
	if err := s.repo.Update(ctx, id, user); err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}

// UpdateUserFields updates specific fields of a user identified by the given UUID.
// It takes a map of field names and their corresponding values to be updated.
// The method interacts with the repository layer to perform the update operation.
// If the update is successful, it returns the updated user object.
// In case of an error during the update, it logs the error and returns it.
func (s *UserService) UpdateUserFields(ctx context.Context, id uuid.UUID, fields map[string]any) (*domain.User, error) {
	updatedUser, err := s.repo.UpdateFields(ctx, id, fields)
	if err != nil {
		log.Printf("Error updating user fields: %v", fields)
		return nil, err
	}

	return updatedUser, nil
}

// DeleteUser removes a user from the repository based on the provided UUID.
// It returns an error if the deletion process fails, logging the error for debugging purposes.
func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	return nil
}

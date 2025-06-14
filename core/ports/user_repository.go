// Package repositories provides functions to interact with the database for managing user data.
// It includes operations such as retrieving all users, fetching a user by ID, creating a new user,
// updating user details, updating specific fields of a user, and deleting a user.
package ports

import (
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/google/uuid"
)

// UserRepository defines the contract for a repository that manages user entities.
// It provides methods for performing CRUD (Create, Read, Update, Delete) operations
// on user data, as well as updating specific fields of a user. The interface abstracts
// the underlying data storage mechanism, allowing for flexibility and easier testing.
type UserRepository interface {
	FindAll() ([]*domain.User, error)
	FindByID(id uuid.UUID) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Save(user *domain.User) error
	Update(id uuid.UUID, user *domain.User) error
	UpdateFields(id uuid.UUID, fields map[string]any) (*domain.User, error)
	Delete(id uuid.UUID) error
}

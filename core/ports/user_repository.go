// Package repositories provides functions to interact with the database for managing user data.
// It includes operations such as retrieving all users, fetching a user by ID, creating a new user,
// updating user details, updating specific fields of a user, and deleting a user.
package ports

import (
	"context"

	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/google/uuid"
)

// UserRepository defines the contract for a repository that manages user entities.
// It provides methods for performing CRUD (Create, Read, Update, Delete) operations
// on user data, as well as updating specific fields of a user. The interface abstracts
// the underlying data storage mechanism, allowing for flexibility and easier testing.
type UserRepository interface {
	FindAll(ctx context.Context) ([]*domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Save(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, id uuid.UUID, user *domain.User) error
	UpdateFields(ctx context.Context, id uuid.UUID, fields map[string]any) (*domain.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

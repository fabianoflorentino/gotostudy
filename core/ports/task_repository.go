// Package repositories provides functions to interact with the database for managing user data.
// It includes operations such as retrieving all users, fetching a user by ID, creating a new user,
// updating user details, updating specific fields of a user, and deleting a user.
package ports

import (
	"context"

	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/google/uuid"
)

// TaskRepository defines the interface for interacting with task data storage.
// It provides methods to find, save, update, and delete tasks.
type TaskRepository interface {
	FindAll(ctx context.Context) ([]*domain.Task, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Task, error)
	Save(ctx context.Context, task *domain.Task) error
	Update(ctx context.Context, id uuid.UUID, task *domain.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
}

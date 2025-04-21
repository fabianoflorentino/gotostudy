// Package services provides the business logic and service layer for the application.
// It acts as an intermediary between the repositories and the controllers,
// handling operations related to users such as retrieval, creation, updating, and deletion.
package ports

import (
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]*domain.User, error)
	GetUserByID(id uuid.UUID) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
	UpdateUser(id uuid.UUID, user *domain.User) (*domain.User, error)
	UpdateUserFields(id uuid.UUID, fields map[string]any) (*domain.User, error)
	DeleteUser(id uuid.UUID) error
}

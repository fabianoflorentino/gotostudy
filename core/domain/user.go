// Package domain contains the core business entities and domain logic
// for the application. It defines the primary data structures and
// relationships that represent the problem domain.
package domain

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user entity in the system.
// It contains the user's unique identifier, username, email, and timestamps
// for when the user was created and last updated. Additionally, it includes
// a list of tasks associated with the user.
type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Tasks     []Task
}

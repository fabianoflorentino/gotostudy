// Package models contains the definitions of data structures (models)
// that represent the core entities in the application. These models
// are used to interact with the database and define the schema for
// the application's data storage.
package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user entity in the system.
// It contains fields for unique identification, username, email, and timestamps.
//
// Fields:
// - ID: A unique identifier for the user, generated as a UUID.
// - Username: The unique username of the user, cannot be null.
// - Email: The unique email address of the user, cannot be null.
// - CreatedAt: The timestamp when the user was created, automatically set.
// - UpdatedAt: The timestamp when the user was last updated, automatically set.
type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

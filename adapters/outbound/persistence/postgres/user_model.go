// Package postgres provides the implementation of data persistence
// for the application using a PostgreSQL database. It defines the
// data models and their relationships, which are used by GORM to
// interact with the database.
package postgres

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user entity in the system.
// It contains fields for storing user-specific information such as
// a unique identifier (ID), username, email, and timestamps for
// creation and last update. The ID is automatically generated as a UUID.
// The Username and Email fields are unique and cannot be null.
// The CreatedAt and UpdatedAt fields are automatically managed by GORM
// to track when the user was created and last updated, respectively.
// The Tasks field establishes a one-to-many relationship with the Task entity,
// where each user can have multiple tasks. Changes to the user will cascade
// to associated tasks on update or delete operations.
type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Tasks     []Task    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

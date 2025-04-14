// Package models contains the definitions of data structures (models)
// that represent the core entities in the application. These models
// are used to interact with the database and define the schema for
// the application's data storage.
package models

import (
	"time"

	"github.com/google/uuid"
)

// Task represents a task entity in the system.
// It includes fields for a unique identifier, name, description, completion status,
// and timestamps for creation and last update.
//
// Fields:
// - ID: A unique identifier for the task, generated as a UUID.
// - Name: The name of the task, which cannot be null.
// - Description: A detailed description of the task, which cannot be null.
// - Completed: A boolean indicating whether the task is completed, defaulting to false.
// - CreatedAt: The timestamp when the task was created, automatically set.
// - UpdatedAt: The timestamp when the task was last updated, automatically set.
type Task struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Completed   bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
}

// Package postgres provides the implementation of data persistence for the application
// using a PostgreSQL database. It defines the data models and their mappings to the
// database schema, enabling seamless interaction with the database.
package postgres

import (
	"time"

	"github.com/google/uuid"
)

// Task represents a task entity in the system. It includes details such as
// the task's unique identifier (ID), title, description, completion status,
// timestamps for creation and updates, and the ID of the user who owns the task.
// The struct is designed to work with GORM for database persistence, with
// annotations specifying primary key, default values, and constraints.
type Task struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Completed   bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
}

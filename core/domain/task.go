// Package domain contains the core business entities and domain models
// used in the application. These models represent the fundamental
// concepts and rules of the problem domain, and they are designed to
// be independent of any specific application layer or infrastructure.
package domain

import (
	"time"

	"github.com/google/uuid"
)

// Task represents a to-do item or activity that a user can create and manage.
// It includes fields for a unique identifier (ID), title, description,
// completion status (Completed), timestamps for creation and updates
// (CreatedAt and UpdatedAt), and the ID of the user who owns the task (UserID).
type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uuid.UUID
}

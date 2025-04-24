package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uuid.UUID
}

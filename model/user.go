package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username string    `gorm:"unique;not null"`
	Email    string    `gorm:"unique;not null"`
}

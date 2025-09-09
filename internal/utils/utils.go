package utils

import (
	"context"
	"errors"
	"regexp"

	"github.com/fabianoflorentino/gotostudy/core"
	"github.com/fabianoflorentino/gotostudy/core/ports"
	"github.com/google/uuid"
)

type EmailChecker interface {
	IsEmailInUse(ctx context.Context, email string, excludeID uuid.UUID) (bool, error)
	IsEmailValid(email string) error
}

// isEmailInUse checks if the given email is already in use by another user in the repository.
// It excludes the user with the specified excludeID from the check.
// Returns true if the email is in use by a different user, false otherwise.
// Returns an error if there is a problem accessing the repository.
func IsEmailInUse(ports ports.UserRepository, ctx context.Context, email string, excludeID uuid.UUID) (bool, error) {
	existingUser, err := ports.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, core.ErrEmailAlreadyExists) {
			return false, nil
		}
		if errors.Is(err, core.ErrUserNotFound) {
			return false, nil
		}
		return false, err
	}

	if excludeID != uuid.Nil && existingUser.ID == excludeID {
		return false, nil
	}

	return true, nil
}

// IsEmailValid validates the format of the given email string.
// It returns core.ErrInvalidEmail if the email format is invalid, or nil if it is valid.
func IsEmailValid(email string) error {
	// Use a simple regex to validate the email format
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(email) {
		return core.ErrInvalidEmail
	}
	return nil
}

// IsTaskTitleValid checks if the provided task title is valid according to predefined rules.
// It returns true if the title is valid, false otherwise.
func IsTaskTitleValid(title string) bool {
	// Task title must be at least 3 characters long and not just whitespace
	if len(title) < 3 {
		return false
	}
	if regexp.MustCompile(`^\s*$`).MatchString(title) {
		return false
	}
	return true
}

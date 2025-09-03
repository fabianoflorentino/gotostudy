package core

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidUpdateField = errors.New("invalid update fields")
	ErrNoTasksFound       = errors.New("no tasks found for user")
	ErrInvalidTaskID      = errors.New("invalid task ID")
	ErrTaskNotFound       = errors.New("task not found")
	ErrInvalidEmail       = errors.New("invalid email format")
)

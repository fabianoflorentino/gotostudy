// Package requests contains the definitions of request structures used for handling
// and validating incoming HTTP requests in the application. These structures
// are typically used to parse and validate JSON payloads from clients.
package requests

// RegisterUserRequest represents the structure of a request payload
// for registering a new user. It includes the user's username and
// email address. Both fields are required, and the email field must
// be a valid email format.
type RegisterUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

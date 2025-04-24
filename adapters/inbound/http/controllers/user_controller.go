// Package controllers provides HTTP handlers for managing user-related operations.
// It defines a UserController struct that interacts with the UserService to handle
// user creation, retrieval, updating, and deletion. The package uses the Gin framework
// for routing and request handling, and it ensures proper validation and error handling
// for incoming requests.
package controllers

import (
	"fmt"
	"net/http"

	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserController is a struct that acts as an HTTP controller for handling
// user-related requests. It depends on the UserService to perform business
// logic operations related to users.
type UserController struct {
	service *services.UserService
}

// NewUserController creates and returns a new instance of UserController.
// It takes a pointer to a UserService as a parameter, which is used to handle
// the business logic related to user operations. This function initializes
// the UserController with the provided service and prepares it for handling
// HTTP requests related to user management.
func NewUserController(s *services.UserService) *UserController {
	return &UserController{service: s}
}

// CreateUser handles the HTTP request for creating a new user.
// It expects a JSON payload containing "username" and "email" fields.
// The "username" field is required, and the "email" field must be a valid email address.
// If the input validation fails, it responds with a 400 Bad Request status and an error message.
// If the user creation process encounters an error, it responds with a 500 Internal Server Error status and an error message.
// On successful user creation, it responds with a 201 Created status and the created user object in the response body.
func (u *UserController) CreateUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.service.RegisterUser(input.Username, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsers handles the HTTP GET request to retrieve all users.
// It interacts with the service layer to fetch the list of users.
// If an error occurs during the retrieval process, it responds with
// an HTTP 500 status code and an error message. Otherwise, it responds
// with an HTTP 200 status code and the list of users in JSON format.
func (u *UserController) GetUsers(c *gin.Context) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID handles the HTTP request to retrieve a user by their unique ID.
// It extracts the user ID from the request parameters, validates it as a UUID,
// and then calls the service layer to fetch the user data. If the ID is invalid,
// it responds with a 400 Bad Request error. If the user is not found, it responds
// with a 404 Not Found error. On success, it returns the user data with a 200 OK status.
func (u *UserController) GetUserByID(c *gin.Context) {
	uid, err := u.parseUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.service.GetUserByID(uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser handles the HTTP request to update an existing user's information.
// It extracts the user ID from the URL parameter, validates the input JSON payload,
// and calls the service layer to update the user details in the system.
// If the user ID is invalid or the input data fails validation, it responds with
// an appropriate HTTP error status and message. On success, it returns the updated
// user information with an HTTP 200 status.
func (u *UserController) UpdateUser(c *gin.Context) {
	uid, err := u.parseUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := u.service.UpdateUser(uid, &domain.User{
		Username: input.Username,
		Email:    input.Email,
	})

	c.JSON(http.StatusOK, user)
}

// UpdateUserFields handles the HTTP request to update specific fields of a user.
// It extracts the user ID from the request parameters, validates it, and parses
// the fields to be updated from the request body. The method ensures that the
// updates are valid before passing them to the service layer for processing.
// If successful, it returns the updated user object in the response. In case of
// errors, appropriate HTTP status codes and error messages are returned.
//
// Parameters:
// - c: The Gin context, which provides request and response handling.
//
// Possible Responses:
//   - HTTP 400: If the user ID is invalid, the update fields are invalid, or
//     there are validation errors.
//   - HTTP 500: If an internal server error occurs during the update process.
//   - HTTP 200: If the user fields are successfully updated, returning the updated user object.
func (u *UserController) UpdateUserFields(c *gin.Context) {
	uid, err := u.parseUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates, err := u.parseUpdateFields(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !u.hasValidUpdates(updates, c) {
		return
	}

	user, err := u.service.UpdateUserFields(uid, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser handles the HTTP DELETE request to remove a user by their unique identifier (UUID).
// It retrieves the user ID from the request parameters, validates it, and attempts to delete the user
// using the service layer. If the UUID is invalid, it responds with a 400 Bad Request status.
// If the user is not found, it responds with a 404 Not Found status. On successful deletion,
// it responds with a 204 No Content status.
func (u *UserController) DeleteUser(c *gin.Context) {
	uid, err := u.parseUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.service.DeleteUser(uid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// parseUUID takes a string representation of a UUID and attempts to parse it into a uuid.UUID object.
// If the provided string is not a valid UUID, it returns an error indicating the issue.
// On success, it returns the parsed uuid.UUID and a nil error.
func (u *UserController) parseUUID(id string) (uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID: %s", err)
	}

	return uid, nil
}

// parseUpdateFields parses and validates the JSON payload from the HTTP request context
// to extract fields for updating a user. It ensures that only valid fields are included
// in the update map. If the JSON payload contains invalid fields or cannot be bound,
// an error is returned.
//
// Parameters:
// - c: The Gin HTTP request context containing the JSON payload.
//
// Returns:
// - A map of valid fields to be updated and their corresponding values.
// - An error if the JSON payload is invalid or contains unsupported fields.
func (u *UserController) parseUpdateFields(c *gin.Context) (map[string]interface{}, error) {
	var updates map[string]interface{}

	if err := c.ShouldBindJSON(&updates); err != nil {
		return nil, err
	}

	validFields := map[string]bool{
		"username": true,
		"email":    true,
	}

	for field := range updates {
		if !validFields[field] {
			return nil, fmt.Errorf("invalid field: %s", field)
		}
	}

	return updates, nil
}

// hasValidUpdates checks if the provided updates map contains any valid fields to update.
// If the map is empty, it responds with a 400 Bad Request status and an error message
// indicating that there are no valid fields to update. Returns true if the updates map
// is not empty, otherwise returns false.
func (u *UserController) hasValidUpdates(updates map[string]interface{}, c *gin.Context) bool {
	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid fields to update"})
		return false
	}

	return true
}

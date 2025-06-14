// Package controllers provides HTTP handlers for managing user-related operations.
// It defines a UserController struct that interacts with the UserService to handle
// user creation, retrieval, updating, and deletion. The package uses the Gin framework
// for routing and request handling, and it ensures proper validation and error handling
// for incoming requests.
package controllers

import (
	"net/http"

	"github.com/fabianoflorentino/gotostudy/adapters/inbound/http/handlers"
	"github.com/fabianoflorentino/gotostudy/adapters/inbound/http/helpers"
	"github.com/fabianoflorentino/gotostudy/adapters/inbound/http/requests"
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/services"
	"github.com/gin-gonic/gin"
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
	var user = &domain.User{}

	if err := handlers.ShouldBindJSON(c, &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, username and email are required"})
		return
	}

	user, err := u.service.RegisterUser(c, user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
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
	users, err := u.service.GetAllUsers(c)
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
	uid, err := helpers.ParseUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if usr := helpers.UserExists(u.service, uid, c); usr != nil {
		c.JSON(http.StatusOK, usr)
	}
}

// UpdateUser handles the HTTP request to update an existing user's information.
// It extracts the user ID from the URL parameter, validates the input JSON payload,
// and calls the service layer to update the user details in the system.
// If the user ID is invalid or the input data fails validation, it responds with
// an appropriate HTTP error status and message. On success, it returns the updated
// user information with an HTTP 200 status.
func (u *UserController) UpdateUser(c *gin.Context) {
	uid, err := helpers.ParseUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if helpers.UserExists(u.service, uid, c) == nil {
		return
	}

	var input requests.RegisterUserRequest
	handlers.ShouldBindJSON(c, &input)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := u.service.UpdateUser(c, uid, &domain.User{
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
	uid, err := helpers.ParseUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updates = handlers.HasValidUpdateUserFields(u.service, c, uid)
	user, err := u.service.UpdateUserFields(c, uid, updates)
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
	uid, err := helpers.ParseUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.service.DeleteUser(c, uid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// File: controllers/user_controller.go
// Description: This file contains the UserController functions.
// It handles the user-related endpoints of the application.
package controllers

import (
	"net/http"

	"github.com/fabianoflorentino/gotostudy/models"
	"github.com/fabianoflorentino/gotostudy/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUsers handles the GET request to retrieve all users.
// It calls the service layer to get the users and returns them as a JSON response.
func GetUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID handles the GET request to retrieve a user by ID.
// It parses the user ID from the URL parameter, calls the service layer to get the user,
// and returns the user as a JSON response.
func GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := services.GetUserByID(parsedUserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if user.ID.String() == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser handles the POST request to create a new user.
// It binds the request body to a User model, calls the service layer to create the user,
// and returns the created user as a JSON response.
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// UpdateUser handles the PUT request to update an existing user.
// It parses the user ID from the URL parameter, binds the request body to a User model,
// calls the service layer to update the user, and returns the updated user as a JSON response.
func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = parsedUserID

	updatedUser, err := services.UpdateUser(parsedUserID, user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// UpdateUserFields handles the PATCH request to update specific fields of a user.
// It parses the user ID from the URL parameter, binds the request body to a map of fields,
// calls the service layer to update the user fields, and returns the updated user as a JSON response.
func UpdateUserFields(c *gin.Context) {
	userID := c.Param("id")

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	var fields map[string]any
	if err := c.ShouldBindJSON(&fields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := services.UpdateUserFields(parsedUserID, fields)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser handles the DELETE request to delete a user by ID.
// It parses the user ID from the URL parameter, calls the service layer to delete the user,
// and returns a success message as a JSON response.
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	err = services.DeleteUser(parsedUserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

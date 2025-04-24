package controllers

import (
	"fmt"
	"net/http"

	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(s *services.UserService) *UserController {
	return &UserController{service: s}
}

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

func (u *UserController) GetUsers(c *gin.Context) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserController) GetUserByID(c *gin.Context) {
	uid := parseUUID(c.Param("id"))

	user, err := u.service.GetUserByID(uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var uid = parseUUID(c.Param("id"))

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

func (u *UserController) UpdateUserFields(c *gin.Context) {
	uid := parseUUID(c.Param("id"))

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

func (u *UserController) DeleteUser(c *gin.Context) {
	uid := parseUUID(c.Param("id"))

	err := u.service.DeleteUser(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func parseUUID(params string) uuid.UUID {
	id, err := uuid.Parse(params)
	if err != nil {
		return uuid.Nil
	}

	return id
}

// private methods

func (u *UserController) parseUpdateFields(c *gin.Context) (map[string]interface{}, error) {
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		return nil, err
	}

	// Validate the fields
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

func (u *UserController) hasValidUpdates(updates map[string]interface{}, c *gin.Context) bool {
	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid fields to update"})
		return false
	}

	return true
}

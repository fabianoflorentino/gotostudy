package controllers

import (
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	Service ports.UserService
}

func New(service ports.UserService) *UserController {
	return &UserController{Service: service}
}

func (h *UserController) RegisterRoutes(r *gin.Engine) {
	r.GET("/users", h.getUsers)
	r.GET("/users/:id", h.getUserByID)
	r.POST("/users", h.createUser)
	r.PUT("/users/:id", h.updateUser)
	r.PATCH("/users/:id", h.updateUserFields)
	r.DELETE("/users/:id", h.deleteUser)
}

func (h *UserController) getUsers(c *gin.Context) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (h *UserController) getUserByID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.Service.GetUserByID(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (h *UserController) createUser(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.Service.CreateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, createdUser)
}

func (h *UserController) updateUser(c *gin.Context) {
	var user domain.User

	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.Service.UpdateUser(userID, &user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, updatedUser)
}

func (h *UserController) updateUserFields(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}

	var fields map[string]any
	if err := c.ShouldBindJSON(&fields); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.Service.UpdateUserFields(userID, fields)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, updatedUser)
}

func (h *UserController) deleteUser(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}

	err = h.Service.DeleteUser(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}

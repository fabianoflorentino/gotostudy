package controllers

import (
	"net/http"

	"github.com/fabianoflorentino/gotostudy/adapters/inbound/http/handlers"
	"github.com/fabianoflorentino/gotostudy/adapters/inbound/http/helpers"
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/services"
	"github.com/gin-gonic/gin"
)

// TaskController handles HTTP requests related to task operations by interacting with the TaskService.
type TaskController struct {
	task *services.TaskService
}

// NewTaskController creates and returns a new instance of TaskController with the provided TaskService.
// It initializes the TaskController's task field with the given TaskService dependency.
func NewTaskController(t *services.TaskService) *TaskController {
	return &TaskController{task: t}
}

// CreateTask handles the HTTP request to create a new task for a specific user.
// It expects a JSON payload with the task details in the request body and a user ID as a URL parameter.
// If the request body is invalid or the user ID is not a valid UUID, it responds with a 400 Bad Request.
// If the task creation fails, it responds with a 422 Unprocessable Entity and the error message.
// On success, it responds with a 201 Created status and the created task in the response body.
func (t *TaskController) CreateTask(c *gin.Context) {

	var task = &domain.Task{}

	if err := handlers.ShouldBindJSON(c, &task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, title and description are required"})
		return
	}

	params, ok := helpers.ValidateUUIDParams(c, "id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
		return
	}

	userID := params[0]

	if err := t.task.CreateTask(c, userID, task); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// FindUserTasks handles HTTP requests to retrieve all tasks associated with a specific user.
// It parses the user ID from the request parameters, validates it, and then fetches the user's tasks.
// If the user ID is invalid, it responds with HTTP 400 Bad Request.
// If an error occurs while retrieving tasks, it responds with HTTP 422 Unprocessable Entity.
// On success, it responds with HTTP 200 OK and the list of tasks.
func (t *TaskController) FindUserTasks(c *gin.Context) {
	params, ok := helpers.ValidateUUIDParams(c, "id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userID := params[0]

	tasks, err := t.task.FindUserTasks(c, userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "user not have tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// FindTaskByID handles HTTP requests to retrieve a specific task by its ID for a given user.
// It expects "id" (user ID) and "task_id" (task ID) as URL parameters.
// If the parameters are invalid UUIDs, it responds with HTTP 400 Bad Request.
// If the task cannot be found or another error occurs, it responds with HTTP 422 Unprocessable Entity.
// On success, it responds with HTTP 200 OK and the task data in JSON format.
func (t *TaskController) FindTaskByID(c *gin.Context) {
	params, ok := helpers.ValidateUUIDParams(c, "id", "task_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user or task ID"})
		return
	}

	userID := params[0]
	taskID := params[1]

	task, err := t.task.FindTaskByID(c, userID, taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask handles HTTP PUT requests to update an existing task for a specific user.
// It parses the user ID and task ID from the URL parameters, binds the request body to a Task struct,
// and calls the service layer to update the task. Returns appropriate HTTP status codes and error messages
// for invalid input or update failures.
func (t *TaskController) UpdateTask(c *gin.Context) {
	var task domain.Task

	params, ok := helpers.ValidateUUIDParams(c, "id", "task_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user or task ID"})
		return
	}

	if err := handlers.ShouldBindJSON(c, &task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, title and description are required"})
		return
	}

	userID := params[0]
	taskID := params[1]

	if err := t.task.UpdateTask(c, userID, taskID, &task); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "")
}

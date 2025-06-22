package services

import (
	"context"

	"github.com/fabianoflorentino/gotostudy/core"
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/ports"
	"github.com/google/uuid"
)

// TaskService provides methods to manage tasks by interacting with the TaskRepository.
// It acts as a service layer between the application logic and the data access layer.
type TaskService struct {
	tsk ports.TaskRepository
}

// NewTaskService creates a new instance of TaskService using the provided TaskRepository.
// It returns a pointer to the initialized TaskService.
func NewTaskService(t ports.TaskRepository) *TaskService {
	return &TaskService{tsk: t}
}

// CreateTask creates a new task for the specified user.
// It first checks if the user exists; if not, it returns core.ErrUserNotFound.
// If the user exists, it attempts to save the task using the underlying task repository.
// Returns an error if saving fails, or nil on success.
func (t *TaskService) CreateTask(ctx context.Context, task *domain.Task) error {
	if !userExists(ctx, task.UserID) {
		return core.ErrUserNotFound
	}

	if err := t.tsk.Save(ctx, task); err != nil {
		return err
	}

	return nil
}

// GetUserTasks retrieves all tasks associated with the specified user ID.
// It accepts a context for request-scoped values and cancellation, and a userID of type uuid.UUID.
// Returns a slice of pointers to domain.Task and an error if the operation fails.
func (t *TaskService) GetUserTasks(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error) {
	if !userExists(ctx, userID) {
		return nil, core.ErrUserNotFound
	}

	// Retrieve tasks for the user from the task repository.
	tasks, err := t.tsk.FindUserTasks(ctx, userID)
	if err != nil {
		return nil, err
	}

	// If no tasks are found, return an error indicating no tasks were found for the user.
	if len(tasks) == 0 {
		return nil, core.ErrNoTasksFound
	}

	return tasks, nil
}

// GetTaskByID retrieves a task by its unique identifier.
// It returns the corresponding Task if found, or an error if the task does not exist,
// the provided taskID is invalid, or another error occurs during retrieval.
//
// Parameters:
//   - ctx: context.Context for controlling cancellation and deadlines.
//   - taskID: uuid.UUID representing the unique identifier of the task.
//
// Returns:
//   - *domain.Task: pointer to the retrieved Task, or nil if not found or on error.
//   - error: error encountered during retrieval, or nil if successful.
func (t *TaskService) GetTaskByID(ctx context.Context, taskID uuid.UUID) (*domain.Task, error) {
	if taskID == uuid.Nil {
		return nil, core.ErrInvalidTaskID
	}

	task, err := t.taskExists(ctx, taskID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// UpdateTask updates an existing task identified by taskID with the provided task details.
// It returns an error if the taskID is invalid, the user does not exist, or if there is a failure
// during the update process.
func (t *TaskService) UpdateTask(ctx context.Context, taskID uuid.UUID, task *domain.Task) error {
	if taskID == uuid.Nil {
		return core.ErrInvalidTaskID
	}

	// Check if the user exists before proceeding with the task update.
	if !userExists(ctx, taskID) {
		return core.ErrUserNotFound
	}

	// Check if the task exists before updating it.
	existingTask, err := t.taskExists(ctx, taskID)
	if err != nil {
		return err
	}

	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.Completed = task.Completed

	if err := t.tsk.Update(ctx, taskID, existingTask); err != nil {
		return err
	}

	return nil
}

// DeleteTask deletes a task identified by the given taskID.
// It returns an error if the taskID is invalid, if the task does not exist,
// or if there is a failure during the deletion process.
func (t *TaskService) DeleteTask(ctx context.Context, taskID uuid.UUID) error {
	// Validate the taskID to ensure it is not a nil UUID.
	if taskID == uuid.Nil {
		return core.ErrInvalidTaskID
	}

	// Check if the task exists before attempting to delete it.'
	if _, err := t.taskExists(ctx, taskID); err != nil {
		return err
	}

	if err := t.tsk.Delete(ctx, taskID); err != nil {
		return err
	}

	return nil
}

// userExists checks if a user with the given userID exists in the system.
// It returns true if the user exists, false otherwise.
func userExists(ctx context.Context, userID uuid.UUID) bool {
	user, err := UserService{}.usr.FindByID(ctx, userID)
	if err != nil {
		return false
	}

	if user == nil {
		return false
	}

	return true
}

// taskExists checks if a task with the given taskID exists in the system.
// It retrieves the task from the repository and returns it if found.
func (t *TaskService) taskExists(ctx context.Context, taskID uuid.UUID) (*domain.Task, error) {
	task, err := t.tsk.FindTaskByID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, core.ErrTaskNotFound
	}

	return task, nil
}

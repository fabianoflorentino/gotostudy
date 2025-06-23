package postgres

import (
	"context"
	"fmt"

	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PostgresTaskRepository is a struct that implements the TaskRepository interface
// for PostgreSQL. It uses GORM for database operations.
type PostgresTaskRepository struct {
	DB *gorm.DB
}

var (
	tasks []Task
	task  Task
)

// NewPostgresTaskRepository creates a new instance of PostgresTaskRepository.
func NewPostgresTaskRepository(db *gorm.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{DB: db}
}

// Save persists the given Task domain entity into the PostgreSQL database.
// It converts the domain.Task to the persistence model and inserts it using GORM.
// Returns an error if the operation fails.
func (t *PostgresTaskRepository) Save(ctx context.Context, task *domain.Task) error {
	model := Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		UserID:      task.UserID,
	}

	return t.DB.Create(&model).Error
}

// FindUserTasks retrieves all tasks associated with the specified user ID from the database.
// It returns a slice of pointers to domain.Task and an error, if any occurs during the query.
//
// Parameters:
//   - ctx: The context for controlling cancellation and timeouts.
//   - userID: The UUID of the user whose tasks are to be retrieved.
//
// Returns:
//   - []*domain.Task: A slice containing pointers to the retrieved tasks.
//   - error: An error object if the operation fails, otherwise nil.
func (t *PostgresTaskRepository) FindUserTasks(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error) {
	if err := t.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}

	tasks := make([]*domain.Task, len(tasks))
	for i, t := range tasks {
		tasks[i] = &domain.Task{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Completed:   t.Completed,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
			UserID:      t.UserID,
		}
	}

	return tasks, nil
}

// FindTaskByID retrieves a task from the database by its unique identifier (taskID).
// It returns a pointer to the domain.Task if found, or an error if the task does not exist
// or if there is a problem accessing the database.
//
// Parameters:
//   - ctx: context.Context for controlling cancellation and deadlines.
//   - taskID: uuid.UUID representing the unique identifier of the task.
//
// Returns:
//   - *domain.Task: pointer to the found task, or nil if not found.
//   - error: error encountered during the operation, or nil if successful.
func (t *PostgresTaskRepository) FindTaskByID(ctx context.Context, taskID uuid.UUID) (*domain.Task, error) {
	if err := t.DB.Where("id = ?", taskID).First(&task).Error; err != nil {
		return nil, err
	}

	return &domain.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		UserID:      task.UserID,
	}, nil
}

// Update updates the task identified by taskID in the PostgreSQL database with the values from tsk.
// It returns an error if the update operation fails.
// ctx is the context for controlling cancellation and timeouts.
// taskID is the unique identifier of the task to be updated.
// tsk is a pointer to the Task domain model containing the updated data.
func (t *PostgresTaskRepository) Update(ctx context.Context, taskID uuid.UUID, tsk *domain.Task) error {
	if err := t.DB.Where("id = ?", taskID).First(&task).Error; err != nil {
		return err
	}

	task.Title = tsk.Title
	task.Description = tsk.Description
	task.Completed = tsk.Completed
	task.UpdatedAt = tsk.UpdatedAt

	return t.DB.Save(&task).Error
}

// UpdateFields updates specific fields of a task identified by taskID in the database.
// The fields parameter is a map where the keys are the names of the fields to update and the values are the new values for those fields.
// Returns the updated Task domain object or an error if the update fails.
func (t *PostgresTaskRepository) UpdateFields(ctx context.Context, taskID uuid.UUID, fields map[string]any) (*domain.Task, error) {
	if err := t.DB.Where("id = ?", taskID).First(&task).Error; err != nil {
		return nil, err
	}

	if _, err := t.hasValidFields(fields); err != nil {
		return nil, err
	}

	task := &domain.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		UserID:      task.UserID,
	}

	return task, nil
}

// Delete removes a task from the database identified by the given taskID.
// It first checks if the task exists, returning an error if not found or if a database error occurs.
// If the task exists, it deletes the task and returns any error encountered during deletion.
func (t *PostgresTaskRepository) Delete(ctx context.Context, taskID uuid.UUID) error {
	if err := t.DB.Where("id = ?", taskID).First(&task).Error; err != nil {
		return err
	}

	return t.DB.Delete(&task).Error
}

// hasValidFields checks if the provided fields map contains only valid task fields.
// It returns true and nil error if at least one valid, non-empty string field is found.
// If an invalid field is encountered, it returns false and an error indicating the invalid field.
// If no valid fields are provided, it returns false and an error.
// Note: This function also performs a database update for the first valid, non-empty string field found.
func (t *PostgresTaskRepository) hasValidFields(fields map[string]any) (bool, error) {
	validFields := map[string]bool{
		"title":       true,
		"description": true,
		"completed":   true,
	}

	for key, value := range fields {
		if !validFields[key] {
			return false, fmt.Errorf("invalid field: %s", key)
		}

		if strValue, ok := value.(string); ok && strValue != "" {
			t.DB.Model(&task).Update(key, strValue)
			return true, nil
		}
	}

	return false, fmt.Errorf("no valid fields provided")
}

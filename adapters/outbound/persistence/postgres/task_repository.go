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

func (t *PostgresTaskRepository) Delete(ctx context.Context, taskID uuid.UUID) error {
	if err := t.DB.Where("id = ?", taskID).First(&task).Error; err != nil {
		return err
	}

	return t.DB.Delete(&task).Error
}

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

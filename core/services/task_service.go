package services

import (
	"context"

	"github.com/fabianoflorentino/gotostudy/core"
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/ports"
	"github.com/google/uuid"
)

type TaskService struct {
	tsk ports.TaskRepository
}

func NewTaskService(t ports.TaskRepository) *TaskService {
	return &TaskService{tsk: t}
}

func (t *TaskService) CreateTask(ctx context.Context, task *domain.Task) error {
	if !userExists(ctx, task.UserID) {
		return core.ErrUserNotFound
	}

	if err := t.tsk.Save(ctx, task); err != nil {
		return err
	}

	return nil
}

func (t *TaskService) GetUserTasks(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error) {
	if !userExists(ctx, userID) {
		return nil, core.ErrUserNotFound
	}

	tasks, err := t.tsk.FindUserTasks(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, core.ErrNoTasksFound
	}

	return tasks, nil
}

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

func (t *TaskService) UpdateTask(ctx context.Context, taskID uuid.UUID, task *domain.Task) error {
	if taskID == uuid.Nil {
		return core.ErrInvalidTaskID
	}

	if !userExists(ctx, taskID) {
		return core.ErrUserNotFound
	}

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

func (t *TaskService) DeleteTask(ctx context.Context, taskID uuid.UUID) error {
	if taskID == uuid.Nil {
		return core.ErrInvalidTaskID
	}

	if _, err := t.taskExists(ctx, taskID); err != nil {
		return err
	}

	if err := t.tsk.Delete(ctx, taskID); err != nil {
		return err
	}

	return nil
}

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

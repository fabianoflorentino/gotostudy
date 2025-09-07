package services

import (
	"context"
	"testing"
	"time"

	"github.com/fabianoflorentino/gotostudy/core"
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/google/uuid"
)

type mockTaskRepository struct {
	tasks map[string]*domain.Task
}

func newMockTaskRepository() *mockTaskRepository {
	return &mockTaskRepository{tasks: make(map[string]*domain.Task)}
}

func (m *mockTaskRepository) Save(ctx context.Context, userID uuid.UUID, task *domain.Task) error {
	m.tasks[task.ID.String()] = task
	return nil
}

func (m *mockTaskRepository) FindUserTasks(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error) {
	var userTasks []*domain.Task
	for _, task := range m.tasks {
		if task.UserID == userID {
			userTasks = append(userTasks, task)
		}
	}

	if len(userTasks) == 0 {
		return nil, core.ErrFindUserTasks
	}

	return userTasks, nil
}

func (m *mockTaskRepository) FindTaskByID(ctx context.Context, userID, taskID uuid.UUID) (*domain.Task, error) {
	task, exists := m.tasks[taskID.String()]
	if !exists {
		return nil, core.ErrTaskNotFound
	}

	return task, nil
}

func (m *mockTaskRepository) Update(ctx context.Context, taskID uuid.UUID, updatedTask *domain.Task) error {
	task, exists := m.tasks[taskID.String()]
	if !exists {
		return core.ErrTaskNotFound
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Completed = updatedTask.Completed
	task.UpdatedAt = updatedTask.UpdatedAt

	m.tasks[taskID.String()] = task
	return nil
}

func (m *mockTaskRepository) Delete(ctx context.Context, taskID uuid.UUID) error {
	_, exists := m.tasks[taskID.String()]
	if !exists {
		return core.ErrTaskNotFound
	}

	delete(m.tasks, taskID.String())
	return nil
}

type mockTaskRepositoryWithError struct{}

func (m *mockTaskRepositoryWithError) Save(ctx context.Context, userID uuid.UUID, task *domain.Task) error {
	return core.ErrCreateTask
}

func (m *mockTaskRepositoryWithError) FindUserTasks(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error) {
	return nil, core.ErrFindUserTasks
}

func (m *mockTaskRepositoryWithError) FindTaskByID(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) (*domain.Task, error) {
	return nil, core.ErrTaskNotFound
}

func (m *mockTaskRepositoryWithError) Update(ctx context.Context, taskID uuid.UUID, updatedTask *domain.Task) error {
	return core.ErrUpdateTask
}

func (m *mockTaskRepositoryWithError) Delete(ctx context.Context, taskID uuid.UUID) error {
	return core.ErrDeleteTask
}

func TestCreateTask(t *testing.T) {
	mockTaskRepo := newMockTaskRepository()
	mockUserRepo := newMockUserRepository()
	taskService := NewTaskService(mockTaskRepo, mockUserRepo)
	userID := uuid.New()

	testNewTask := []struct {
		Context context.Context
		Task    domain.Task
	}{
		{context.Background(), domain.Task{ID: uuid.New(), UserID: userID, Title: "Test Task 1", Description: "This is a test task 1", Completed: false}},
		{context.Background(), domain.Task{ID: uuid.New(), UserID: userID, Title: "Test Task 2", Description: "This is a test task 2", Completed: true}},
		{context.Background(), domain.Task{ID: uuid.New(), UserID: userID, Title: "Test Task 3", Description: "This is a test task 3", Completed: false}},
	}

	t.Run("CreateTask", func(t *testing.T) {
		for _, task := range testNewTask {
			t.Run(task.Task.Title, func(t *testing.T) {
				createTask, _ := taskService.CreateTask(task.Context, userID, &task.Task)
				if createTask != uuid.Nil {
					t.Errorf("expected task to be created, got nil")
				}
			})
		}
	})

	t.Run("CreateTask_DuplicateID", func(t *testing.T) {
		mockUserRepo.users[userID.String()] = &domain.User{
			ID:       userID,
			Username: "testuser",
			Email:    "testuser@example.com",
		}

		task := domain.Task{ID: uuid.New(), UserID: userID, Title: "Unique Task", Description: "This is a unique task", Completed: false}

		_, err := taskService.CreateTask(context.Background(), userID, &task)
		if err != core.ErrCreateTask {
			t.Fatalf("Expected ErrCreateTask for duplicate ID, got: %v", err)
		}

		_, err = taskService.CreateTask(context.Background(), userID, &task)
		if err != core.ErrCreateTask {
			t.Errorf("Expected ErrCreateTask for duplicate ID, got: %v", err)
		}
	})

	t.Run("CreateTaskWithError", func(t *testing.T) {
		mockTaskRepoWithError := &mockTaskRepositoryWithError{}
		taskServiceWithError := NewTaskService(mockTaskRepoWithError, mockUserRepo)
		task := domain.Task{ID: uuid.New(), UserID: userID, Title: "Error Task", Description: "This should fail", Completed: false}
		// Ensure user exists in mockUserRepo to avoid user not found error
		mockUserRepo.users[userID.String()] = &domain.User{
			ID:        userID,
			Username:  "testuser",
			Email:     "testuser@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		_, err := taskServiceWithError.CreateTask(context.Background(), userID, &task)
		if err != core.ErrCreateTask {
			t.Errorf("Expected ErrCreateTask, got: %v", err)
		}
	})
}

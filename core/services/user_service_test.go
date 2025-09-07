package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/fabianoflorentino/gotostudy/core"
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/internal/utils"
	"github.com/google/uuid"
)

// mockUserRepository is a mock implementation of UserRepository for testing
type mockUserRepository struct {
	users map[string]*domain.User
}

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (m *mockUserRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	for _, user := range m.users {
		users = append(users, user)
	}
	return users, nil
}

func (m *mockUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, core.ErrUserNotFound
}

func (m *mockUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	if email == "error@example.com" {
		return nil, fmt.Errorf("simulated error for email checker")
	}
	if user, exists := m.users[email]; exists {
		return user, nil
	}

	return nil, core.ErrUserNotFound
}

func (m *mockUserRepository) Save(ctx context.Context, user *domain.User) error {
	if user.Email == "save_error@example.com" {
		return fmt.Errorf("simulated save error")
	}
	m.users[user.Email] = user

	return nil
}

func (m *mockUserRepository) Update(ctx context.Context, id uuid.UUID, user *domain.User) error {
	for email, existingUser := range m.users {
		if existingUser.ID == id {
			delete(m.users, email)
			m.users[user.Email] = user
			return nil
		}
	}

	return core.ErrUserNotFound
}

func (m *mockUserRepository) UpdateFields(ctx context.Context, id uuid.UUID, fields map[string]any) (*domain.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			// Simple field update - in a real implementation this would be more robust
			if email, ok := fields["email"].(string); ok {
				user.Email = email
			}
			if username, ok := fields["username"].(string); ok {
				user.Username = username
			}
			return user, nil
		}
	}
	return nil, core.ErrUserNotFound
}

func (m *mockUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	for email, user := range m.users {
		if user.ID == id {
			delete(m.users, email)
			return nil
		}
	}
	return nil
}

type mockUserRepositoryWithError struct{}

func (m *mockUserRepositoryWithError) FindAll(ctx context.Context) ([]*domain.User, error) {
	return nil, core.ErrFindAllUsers
}

func (m *mockUserRepositoryWithError) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return nil, core.ErrUserNotFound
}

func (m *mockUserRepositoryWithError) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, core.ErrFindByEmail
}

func (m *mockUserRepositoryWithError) Save(ctx context.Context, user *domain.User) error {
	return core.ErrSaveUser
}

func (m *mockUserRepositoryWithError) Update(ctx context.Context, id uuid.UUID, user *domain.User) error {
	return core.ErrUserNotFound
}

func (m *mockUserRepositoryWithError) UpdateFields(ctx context.Context, id uuid.UUID, fields map[string]any) (*domain.User, error) {
	return nil, core.ErrUserNotFound
}

func (m *mockUserRepositoryWithError) Delete(ctx context.Context, id uuid.UUID) error {
	return core.ErrDeleteUser
}

func TestRegisterUser(t *testing.T) {
	repo := newMockUserRepository()
	service := NewUserService(repo)

	testNewUsers := []struct {
		Context context.Context
		User    domain.User
	}{
		{Context: context.Background(), User: domain.User{Username: "testuser", Email: "test@example.com"}},
		{Context: context.Background(), User: domain.User{Username: "testuser2", Email: "test2@example.com"}},
	}

	t.Run("RegisterUser", func(t *testing.T) {
		for _, user := range testNewUsers {
			t.Run(user.User.Username, func(t *testing.T) {
				createdUser, err := service.RegisterUser(user.Context, &user.User)
				if err != nil {
					t.Fatalf("Failed to create user: %v", err)
				}
				if createdUser.Username != user.User.Username || createdUser.Email != user.User.Email {
					t.Errorf("Created user does not match input: got %+v, want %+v",
						createdUser, user.User)
				}
			})
		}
	})

	t.Run("RegisterUser_DuplicateEmail", func(t *testing.T) {
		user := domain.User{Username: "duplicateuser", Email: "duplicate@example.com"}
		_, err := service.RegisterUser(context.Background(), &user)
		if err != nil {
			t.Fatalf("Failed to create user: %v", err)
		}

		// Try to create the user again
		_, err = service.RegisterUser(context.Background(), &user)
		if err == nil {
			t.Errorf("Expected error when creating duplicate user, got nil")
		}
	})

	t.Run("RegisterUser_InvalidEmail", func(t *testing.T) {
		user := domain.User{Username: "invalidUser", Email: "invalidemail"}
		_, err := service.RegisterUser(context.Background(), &user)
		if err == nil {
			t.Errorf("Expected error when creating user with invalid email, got nil")
		}
	})

	t.Run("RegisterUser_EmailInUse", func(t *testing.T) {
		user := domain.User{Username: "emailInUseUser", Email: "error@example.com"}
		_, err := service.RegisterUser(context.Background(), &user)
		if err == nil {
			t.Errorf("Expected error when creating user with email in use, got nil")
		}
	})

	t.Run("RegisterUser_EmailInUseError", func(t *testing.T) {
		user := domain.User{Username: "errorUser", Email: "error@example.com"}
		_, emailCheckError := utils.IsEmailInUse(repo, context.Background(), user.Email, uuid.Nil)
		if emailCheckError == nil {
			t.Fatalf("Expected error from email checker, got nil")
		}

		_, err := service.RegisterUser(context.Background(), &user)
		if err == nil {
			t.Errorf("Expected error when email check fails, got nil")
		}
	})

	t.Run("RegisterUser_SaveError", func(t *testing.T) {
		log.SetOutput(io.Discard)

		user := domain.User{Username: "errorUser", Email: "save_error@example.com"}
		_, err := service.RegisterUser(context.Background(), &user)
		if err == nil {
			t.Errorf("Expected error when saving user fails, got nil")
		}
	})
}

func TestGetAllUsers(t *testing.T) {
	repo := newMockUserRepository()
	service := NewUserService(repo)

	// Create some test users
	testUsers := []domain.User{
		{Username: "testuser1", Email: "test1@example.com"},
		{Username: "testuser2", Email: "test2@example.com"},
		{Username: "testuser3", Email: "test3@example.com"},
	}

	for _, user := range testUsers {
		_, err := service.RegisterUser(context.Background(), &user)
		if err != nil {
			t.Fatalf("Failed to create user: %v", err)
		}
	}

	t.Run("GetAllUsers", func(t *testing.T) {
		users, err := service.GetAllUsers(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all users: %v", err)
		}
		if len(users) != 3 {
			t.Errorf("Expected 3 users, got %d", len(users))
		}
	})

	t.Run("GetAllUsers_Empty", func(t *testing.T) {
		emptyRepo := newMockUserRepository()
		emptyService := NewUserService(emptyRepo)

		users, err := emptyService.GetAllUsers(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all users from empty repo: %v", err)
		}
		if len(users) != 0 {
			t.Errorf("Expected 0 users, got %d", len(users))
		}
	})

	t.Run("GetAllUsers_Error", func(t *testing.T) {
		errorRepo := &mockUserRepositoryWithError{}
		errorService := NewUserService(errorRepo)

		_, err := errorService.GetAllUsers(context.Background())
		if err == nil {
			t.Errorf("Expected error when repository fails, got nil")
		}
	})
}

func TestGetUserByID(t *testing.T) {
	repo := newMockUserRepository()
	service := NewUserService(repo)

	// Create a test user
	user := domain.User{Username: "testuser", Email: "testuser@example.com"}
	_, err := service.RegisterUser(context.Background(), &user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	t.Run("GetUserByID", func(t *testing.T) {
		user, err := service.GetUserByID(context.Background(), user.ID)
		if err != nil {
			t.Fatalf("Failed to get user by ID: %v", err)
		}
		if user == nil {
			t.Errorf("Expected user, got nil")
		}
	})

	t.Run("GetUserByID_NotFound", func(t *testing.T) {
		nonExistentID := uuid.New()
		user, err := service.GetUserByID(context.Background(), nonExistentID)
		if !errors.Is(err, core.ErrUserNotFound) {
			t.Fatalf("Expected ErrUserNotFound, got: %v", err)
		}
		if user != nil {
			t.Errorf("Expected nil for non-existent user, got %+v", user)
		}
	})

	t.Run("GetUserByID_Error", func(t *testing.T) {
		errorRepo := &mockUserRepositoryWithError{}
		errorService := NewUserService(errorRepo)

		_, err := errorService.GetUserByID(context.Background(), user.ID)
		if err == nil {
			t.Errorf("Expected error when repository fails, got nil")
		}
	})

	t.Run("GetUserByID_InvalidID", func(t *testing.T) {
		invalidID := uuid.Nil
		_, err := service.GetUserByID(context.Background(), invalidID)
		if err == nil {
			t.Errorf("Expected error when using invalid ID, got nil")
		}
	})
}

func TestUpdateUser(t *testing.T) {
	repo := newMockUserRepository()
	service := NewUserService(repo)

	// Create a test user
	user := domain.User{ID: uuid.New(), Username: "testuser", Email: "testuser@example.com"}
	_, err := service.RegisterUser(context.Background(), &user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	t.Run("UpdateUser", func(t *testing.T) {
		updatedUser := domain.User{ID: user.ID, Username: "updateduser", Email: "updateduser@example.com"}
		err := service.UpdateUser(context.Background(), updatedUser.ID, &updatedUser)
		if err != nil {
			t.Fatalf("Failed to update user: %v", err)
		}
	})

	t.Run("UpdateUser_NotFound", func(t *testing.T) {
		log.SetOutput(io.Discard)

		nonExistentID := uuid.New()
		updatedUser := domain.User{ID: nonExistentID, Username: "updateduser", Email: "updateduser_notfound@example.com"}
		err := service.UpdateUser(context.Background(), updatedUser.ID, &updatedUser)
		if !errors.Is(err, core.ErrUserNotFound) {
			t.Fatalf("Expected ErrUserNotFound, got: %v", err)
		}
	})

	t.Run("UpdateUser_Error", func(t *testing.T) {
		errorRepo := &mockUserRepositoryWithError{}
		errorService := NewUserService(errorRepo)

		err := errorService.UpdateUser(context.Background(), user.ID, &user)
		if err == nil {
			t.Errorf("Expected error when repository fails, got nil")
		}
	})

	t.Run("UpdateUser_InvalidEmail", func(t *testing.T) {
		invalidEmailUser := domain.User{ID: user.ID, Username: "invalidemailuser", Email: "invalidemail"}
		err := service.UpdateUser(context.Background(), invalidEmailUser.ID, &invalidEmailUser)
		if err == nil {
			t.Errorf("Expected error when updating user with invalid email, got nil")
		}
	})

	t.Run("UpdateUser_AlreadyExists", func(t *testing.T) {
		// Create another user to cause email conflict
		anotherUser := domain.User{Username: "anotheruser", Email: "anotheruser@example.com"}
		_, err := service.RegisterUser(context.Background(), &anotherUser)
		if err != nil {
			t.Fatalf("Failed to create another user: %v", err)
		}

		// Try to update the original user to have the same email as the new user
		updatedUser := domain.User{ID: user.ID, Username: "updateduser", Email: anotherUser.Email}
		err = service.UpdateUser(context.Background(), updatedUser.ID, &updatedUser)
		if !errors.Is(err, core.ErrEmailAlreadyExists) {
			t.Fatalf("Expected ErrEmailAlreadyExists, got: %v", err)
		}
	})
}

func TestUpdateUserFields(t *testing.T) {
	repo := newMockUserRepository()
	service := NewUserService(repo)

	// Create a test user
	user := domain.User{ID: uuid.New(), Username: "testuser", Email: "testuser@example.com"}
	_, err := service.RegisterUser(context.Background(), &user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	t.Run("UpdateUserFields", func(t *testing.T) {
		updatedFields := map[string]interface{}{
			"username": "updateduser",
			"email":    "updateduser@example.com",
		}
		_, err := service.UpdateUserFields(context.Background(), user.ID, updatedFields)
		if err != nil {
			t.Fatalf("Failed to update user fields: %v", err)
		}
	})

	t.Run("UpdateUserFields_NotFound", func(t *testing.T) {
		log.SetOutput(io.Discard)

		nonExistentID := uuid.New()
		updatedFields := map[string]interface{}{
			"username": "updateduser",
			"email":    "updateduser@example.com",
		}
		_, err := service.UpdateUserFields(context.Background(), nonExistentID, updatedFields)
		if !errors.Is(err, core.ErrUserNotFound) {
			t.Fatalf("Expected ErrUserNotFound, got: %v", err)
		}
	})

	t.Run("UpdateUserFields_Error", func(t *testing.T) {
		nonExistentID := uuid.New()
		updatedFields := map[string]any{
			"username": "updateduser",
			"email":    "updateduser@example.com",
		}
		_, err := service.UpdateUserFields(context.Background(), nonExistentID, updatedFields)
		if !errors.Is(err, core.ErrUserNotFound) {
			t.Fatalf("Expected ErrUserNotFound, got: %v", err)
		}
	})

	t.Run("UpdateUserFields_Error", func(t *testing.T) {
		errorRepo := &mockUserRepositoryWithError{}
		errorService := NewUserService(errorRepo)

		updatedFields := map[string]interface{}{
			"username": "updateduser",
			"email":    "updateduser@example.com",
		}
		_, err := errorService.UpdateUserFields(context.Background(), user.ID, updatedFields)
		if err == nil {
			t.Errorf("Expected error when repository fails, got nil")
		}
	})

	t.Run("UpdateUserFields_InvalidEmail", func(t *testing.T) {
		updatedFields := map[string]any{
			"email": "invalidemail",
		}
		_, err := service.UpdateUserFields(context.Background(), user.ID, updatedFields)
		if err == nil {
			t.Errorf("Expected error when updating user with invalid email, got nil")
		}
	})

	t.Run("UpdateUserFields_AlreadyExists", func(t *testing.T) {
		// Create another user to cause email conflict
		anotherUser := domain.User{Username: "anotheruser", Email: "anotheruser@example.com"}
		_, err := service.RegisterUser(context.Background(), &anotherUser)
		if err != nil {
			t.Fatalf("Failed to create another user: %v", err)
		}

		// Try to update the original user to have the same email as the new user
		updatedFields := map[string]interface{}{
			"email": anotherUser.Email,
		}
		_, err = service.UpdateUserFields(context.Background(), user.ID, updatedFields)
		if !errors.Is(err, core.ErrEmailAlreadyExists) {
			t.Fatalf("Expected ErrEmailAlreadyExists, got: %v", err)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	repo := newMockUserRepository()
	service := NewUserService(repo)

	// Create a test user
	user := domain.User{ID: uuid.New(), Username: "testuser", Email: "testuser@example.com"}
	_, err := service.RegisterUser(context.Background(), &user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	t.Run("DeleteUser", func(t *testing.T) {
		err := service.DeleteUser(context.Background(), user.ID)
		if err != nil {
			t.Fatalf("Failed to delete user: %v", err)
		}
	})

	t.Run("DeleteUser_NotFound", func(t *testing.T) {
		log.SetOutput(io.Discard)

		nonExistentID := uuid.New()
		err := service.DeleteUser(context.Background(), nonExistentID)
		if err != nil {
			t.Fatalf("Expected no error when deleting non-existent user, got: %v", err)
		}
	})

	t.Run("DeleteUser_Error", func(t *testing.T) {
		errorRepo := &mockUserRepositoryWithError{}
		errorService := NewUserService(errorRepo)

		err := errorService.DeleteUser(context.Background(), user.ID)
		if err == nil {
			t.Errorf("Expected error when repository fails, got nil")
		}
	})
}

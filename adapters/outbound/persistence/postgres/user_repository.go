// Package postgres provides an implementation of the UserRepository interface
// using a PostgreSQL database as the persistence layer. It leverages the GORM
// library to interact with the database and provides methods for CRUD operations
// on user entities, including saving, retrieving, updating, and deleting users.
package postgres

import (
	"context"
	"fmt"

	"github.com/fabianoflorentino/gotostudy/core"
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PostgresUserRepository is a struct that provides methods to interact with the
// PostgreSQL database for user-related operations. It uses GORM as the ORM
// (Object-Relational Mapping) library to manage database interactions.
// The DB field is a pointer to a GORM database connection instance.
type PostgresUserRepository struct {
	DB *gorm.DB
}

// models is a slice of User structs, representing a collection of user data
// that can be used for operations such as querying or processing multiple users.
var (
	models []User
	model  User
)

// NewPostgresUserRepository creates a new instance of PostgresUserRepository,
// which implements the UserRepository interface. It takes a gorm.DB instance
// as a parameter to interact with the PostgreSQL database and returns the
// repository implementation. This function is typically used to initialize
// the repository layer for user-related database operations.
func NewPostgresUserRepository(db *gorm.DB) ports.UserRepository {
	return &PostgresUserRepository{DB: db}
}

// Save persists a given User entity into the PostgreSQL database.
// It converts the domain.User object into a database model (User)
// and saves it using the GORM ORM. The method also processes the
// associated tasks of the user, preparing them for persistence.
// Returns an error if the operation fails.
func (r *PostgresUserRepository) Save(ctx context.Context, user *domain.User) error {
	tasks := make([]*domain.Task, len(user.Tasks))
	for i, task := range user.Tasks {
		tasks[i] = &domain.Task{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			UserID:      task.UserID,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		}
	}
	model := User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return r.DB.Create(&model).Error
}

// FindAll retrieves all user records from the database and converts them
// into a slice of domain.User objects. It queries the database using the
// underlying GORM instance (r.DB) and maps the retrieved data to the domain
// layer. If an error occurs during the database query, it returns the error.
// Otherwise, it returns a slice of pointers to domain.User objects.
func (r *PostgresUserRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	if err := r.DB.Find(&models).Error; err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(models))
	for i, model := range models {
		users[i] = &domain.User{
			ID:        model.ID,
			Username:  model.Username,
			Email:     model.Email,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
			Tasks:     nil,
		}
	}

	return users, nil
}

// FindByID retrieves a user from the PostgreSQL database by their unique identifier (UUID).
// It returns a pointer to the User domain object if found, or an error if the user does not exist
// or if there is an issue with the database query.
func (r *PostgresUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	tasks := make([]domain.Task, len(model.Tasks))

	if err := r.DB.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, core.ErrUserNotFound
	}

	for i, task := range model.Tasks {
		tasks[i] = domain.Task{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			UserID:      task.UserID,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		}
	}

	return &domain.User{
		ID:        model.ID,
		Username:  model.Username,
		Email:     model.Email,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Tasks:     tasks,
	}, nil
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var model User

	if err := r.DB.Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       model.ID,
		Username: model.Username,
		Email:    model.Email,
		Tasks:    nil,
	}, nil
}

// Update updates an existing user record in the database with the provided user details.
// It retrieves the user by the given UUID, updates the fields (Username and Email),
// and saves the changes back to the database. If any error occurs during the process,
// it returns the error.
func (r *PostgresUserRepository) Update(ctx context.Context, id uuid.UUID, user *domain.User) error {
	if err := r.DB.Where("id = ?", id).First(&model).Error; err != nil {
		return err
	}

	model.Username = user.Username
	model.Email = user.Email
	model.UpdatedAt = user.UpdatedAt

	return r.DB.Save(&model).Error
}

// UpdateFields updates specific fields of a user in the database identified by the given UUID.
// It accepts a map of field names and their new values, and applies the updates to the user record.
// If the "username" or "email" fields are present in the map, they are updated accordingly.
// The method retrieves the user record, updates the specified fields, and saves the changes back to the database.
// Returns the updated user as a domain.User object or an error if the operation fails.
func (r *PostgresUserRepository) UpdateFields(ctx context.Context, id uuid.UUID, fields map[string]any) (*domain.User, error) {
	if err := r.DB.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}

	if _, err := r.hasValidFields(fields); err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:        model.ID,
		Username:  model.Username,
		Email:     model.Email,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Tasks:     nil,
	}

	return user, nil
}

// Delete removes a user record from the database based on the provided UUID.
// It first attempts to retrieve the user record with the given ID to ensure it exists.
// If the record is found, it deletes the record from the database.
// Returns an error if the record is not found or if any database operation fails.
func (r *PostgresUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.DB.Where("id = ?", id).First(&model).Error; err != nil {
		return err
	}

	return r.DB.Delete(&model).Error
}

func (r *PostgresUserRepository) hasValidFields(fields map[string]any) (bool, error) {
	validFields := map[string]bool{
		"username": true,
		"email":    true,
	}

	for key, value := range fields {
		if !validFields[key] {
			return false, fmt.Errorf("invalid field: %s", key)
		}

		if strValue, ok := value.(string); ok && strValue != "" {
			r.DB.Model(&model).Update(key, strValue)
			return true, nil
		}
	}

	return false, fmt.Errorf("no valid fields provided")
}

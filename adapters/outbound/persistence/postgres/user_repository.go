package postgres

import (
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	DB *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) ports.UserRepository {
	return &PostgresUserRepository{DB: db}
}

func (r *PostgresUserRepository) Save(user *domain.User) error {
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
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return r.DB.Create(&model).Error
}

func (r *PostgresUserRepository) FindAll() ([]*domain.User, error) {
	var models []User
	if err := r.DB.Find(&models).Error; err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(models))
	for i, model := range models {
		users[i] = &domain.User{
			ID:       model.ID,
			Username: model.Username,
			Email:    model.Email,
			Tasks:    nil,
		}
	}

	return users, nil
}

func (r *PostgresUserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	var model User
	tasks := make([]domain.Task, len(model.Tasks))

	if err := r.DB.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
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
		ID:       model.ID,
		Username: model.Username,
		Email:    model.Email,
		Tasks:    tasks,
	}, nil
}

func (r *PostgresUserRepository) Update(id uuid.UUID, user *domain.User) error {
	var model User
	if err := r.DB.First(&model, "id = ?", id).Error; err != nil {
		return err
	}

	model.Username = user.Username
	model.Email = user.Email

	return r.DB.Save(&model).Error
}

func (r *PostgresUserRepository) UpdateFields(id uuid.UUID, fields map[string]any) (*domain.User, error) {
	var model User
	if err := r.DB.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if username, ok := fields["username"]; ok {
		model.Username = username.(string)
	}
	if email, ok := fields["email"]; ok {
		model.Email = email.(string)
	}

	if err := r.DB.Save(&model).Error; err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:       model.ID,
		Username: model.Username,
		Email:    model.Email,
		Tasks:    nil,
	}

	return user, nil
}

func (r *PostgresUserRepository) Delete(id uuid.UUID) error {
	var model User

	if err := r.DB.First(&model, "id = ?", id).Error; err != nil {
		return err
	}

	return r.DB.Delete(&model).Error
}

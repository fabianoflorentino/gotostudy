package services

import (
	"log"

	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/ports"
	"github.com/google/uuid"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(r ports.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) RegisterUser(name, email string) (*domain.User, error) {
	user := &domain.User{
		ID:       uuid.New(),
		Username: name,
		Email:    email,
		Tasks:    []domain.Task{},
	}

	if err := s.repo.Save(user); err != nil {
		log.Printf("Error saving user: %v", err)
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetUserByID(id uuid.UUID) (*domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		log.Printf("Error fetching user by ID: %v", err)
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(id uuid.UUID, user *domain.User) error {
	if err := s.repo.Update(id, user); err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}

func (s *UserService) UpdateUserFields(id uuid.UUID, fields map[string]any) (*domain.User, error) {
	updatedUser, err := s.repo.UpdateFields(id, fields)
	if err != nil {
		log.Printf("Error updating user fields: %v", err)
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
	if err := s.repo.Delete(id); err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	return nil
}

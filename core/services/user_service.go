package services

import (
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/ports"
	"github.com/google/uuid"
)

type UserServiceImplementation struct {
	user ports.UserRepository
}

func New(user ports.UserRepository) ports.UserService {
	return &UserServiceImplementation{user: user}
}

func (u *UserServiceImplementation) GetAllUsers() ([]*domain.User, error) {
	users, err := u.user.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserServiceImplementation) GetUserByID(id uuid.UUID) (*domain.User, error) {
	user, err := u.user.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImplementation) CreateUser(user *domain.User) (*domain.User, error) {
	createdUser, err := u.user.Save(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *UserServiceImplementation) UpdateUser(id uuid.UUID, user *domain.User) (*domain.User, error) {
	updatedUser, err := u.user.Update(id, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
func (u *UserServiceImplementation) UpdateUserFields(id uuid.UUID, fields map[string]any) (*domain.User, error) {
	updatedUser, err := u.user.UpdateFields(id, fields)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *UserServiceImplementation) DeleteUser(id uuid.UUID) error {
	err := u.user.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

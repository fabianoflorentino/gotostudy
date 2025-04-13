package services

import (
	"github.com/fabianoflorentino/gotostudy/models"
	"github.com/fabianoflorentino/gotostudy/repositories"
	"github.com/google/uuid"
)

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func GetUserByID(id uuid.UUID) (models.User, error) {
	return repositories.GetUserByID(id)
}

func CreateUser(user models.User) (models.User, error) {
	return repositories.CreateUser(user)
}

func UpdateUser(id uuid.UUID, user models.User) (models.User, error) {
	return repositories.UpdateUser(id, user)
}

func UpdateUserFields(id uuid.UUID, fields map[string]any) (models.User, error) {
	return repositories.UpdateUserFields(id, fields)
}

func DeleteUser(id uuid.UUID) error {
	return repositories.DeleteUser(id)
}

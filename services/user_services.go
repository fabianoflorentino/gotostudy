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

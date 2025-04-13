package services

import (
	"github.com/fabianoflorentino/gotostudy/models"
	"github.com/fabianoflorentino/gotostudy/repositories"
)

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

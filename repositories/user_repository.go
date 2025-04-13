package repositories

import (
	"github.com/fabianoflorentino/gotostudy/database"
	"github.com/fabianoflorentino/gotostudy/models"
)

var (
	users []models.User
)

func GetAllUsers() ([]models.User, error) {
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

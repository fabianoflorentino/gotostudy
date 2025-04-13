package repositories

import (
	"github.com/fabianoflorentino/gotostudy/database"
	"github.com/fabianoflorentino/gotostudy/models"
	"github.com/google/uuid"
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

func GetUserByID(id uuid.UUID) (models.User, error) {
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	if err := database.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

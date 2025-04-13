package repositories

import (
	"github.com/fabianoflorentino/gotostudy/database"
	"github.com/fabianoflorentino/gotostudy/models"
	"github.com/google/uuid"
)

var (
	users []models.User
	user  models.User
)

func GetAllUsers() ([]models.User, error) {
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uuid.UUID) (models.User, error) {
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

func UpdateUser(id uuid.UUID, user models.User) (models.User, error) {
	if err := database.DB.Save(&user).Where("id = ?", id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func UpdateUserFields(id uuid.UUID, fields map[string]any) (models.User, error) {
	if err := database.DB.Model(&user).Where("id = ?", id).Updates(fields).Error; err != nil {
		return models.User{}, err
	}

	if err := database.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func DeleteUser(id uuid.UUID) error {
	if err := database.DB.First(&user, id).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

package persistence

import (
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SqlUserRepository struct {
	DB *gorm.DB
}

var (
	users []*domain.User
	user  *domain.User
)

func New(db *gorm.DB) ports.UserRepository {
	return &SqlUserRepository{DB: db}
}

func (r *SqlUserRepository) FindAll() ([]*domain.User, error) {
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *SqlUserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *SqlUserRepository) Save(user *domain.User) (*domain.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *SqlUserRepository) Update(id uuid.UUID, user *domain.User) (*domain.User, error) {
	if err := r.DB.Model(&domain.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *SqlUserRepository) UpdateFields(id uuid.UUID, fields map[string]any) (*domain.User, error) {
	if err := r.DB.Model(&domain.User{}).Where("id = ?", id).Updates(fields).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (r *SqlUserRepository) Delete(id uuid.UUID) error {
	if err := r.DB.Delete(&domain.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

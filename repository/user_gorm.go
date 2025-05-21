package repository

import (
	"github.com/KITTTPOB-bank/hospitalapi/initializers"
	"github.com/KITTTPOB-bank/hospitalapi/models"
)

type userRepositoryGorm struct{}

func NewUserRepository() UserRepository {
	return &userRepositoryGorm{}
}

func (r *userRepositoryGorm) FindByUsername(username, hospital string) (*models.User, error) {
	var user models.User
	result := initializers.DB.Where("username = ? AND hospital = ?", username, hospital).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepositoryGorm) Create(user *models.User) error {
	return initializers.DB.Create(user).Error
}

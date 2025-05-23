package repository

import (
	"fmt"

	"github.com/KITTTPOB-bank/hospitalapi/initializers"
	"github.com/KITTTPOB-bank/hospitalapi/models"
)

type userRepositoryGorm struct{}

func NewUserRepository() UserRepository {
	return &userRepositoryGorm{}
}

func (r *userRepositoryGorm) FindByUsername(username, hospital string) (*models.User, error) {
	var user *models.User

	result := initializers.DB.Raw("SELECT * FROM staff WHERE username = ? AND hospital = ?", username, hospital).Scan(&user)
	fmt.Println(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *userRepositoryGorm) GenerateStaff(user *models.User) error {
	return initializers.DB.Create(user).Error
}

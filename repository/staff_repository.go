package repository

import "github.com/KITTTPOB-bank/hospitalapi/models"

type UserRepository interface {
	FindByUsername(username, hospital string) (*models.User, error)
	GenerateStaff(user *models.User) error
}

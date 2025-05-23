package usecase

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/KITTTPOB-bank/hospitalapi/initializers"
	"github.com/KITTTPOB-bank/hospitalapi/models"
	"github.com/KITTTPOB-bank/hospitalapi/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

type UserUsecase interface {
	Register(username, password, hospital string) (string, error)
	Login(username, password, hospital string) (string, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{repo: r}
}

func (u *userUsecase) Register(username, password, hospital string) (string, error) {
	existing, _ := u.repo.FindByUsername(username, hospital)
	fmt.Println(existing)
	if existing != nil {
		return "", errors.New("username already used")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	now := time.Now()
	user := &models.User{
		Username:  username,
		Hospital:  hospital,
		Password:  string(passwordHash),
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = u.repo.GenerateStaff(user)
	if err != nil {
		return "", err
	}

	token, err := generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userUsecase) Login(username, password, hospital string) (string, error) {
	user, err := u.repo.FindByUsername(username, hospital)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid Password")
	}

	token, err := generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"hospital": user.Hospital,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_KEY")

	return token.SignedString([]byte(secret))
}

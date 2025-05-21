package controller

import (
	"net/http"

	"github.com/KITTTPOB-bank/hospitalapi/models"
	"github.com/KITTTPOB-bank/hospitalapi/repository"
	"github.com/KITTTPOB-bank/hospitalapi/usecase"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input models.Login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRepo := repository.NewUserRepository()
	userUC := usecase.NewUserUsecase(userRepo)

	user, err := userUC.Register(input.Username, input.Password, input.Hospital)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

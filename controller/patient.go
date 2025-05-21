package controller

import (
	"net/http"

	"github.com/KITTTPOB-bank/hospitalapi/models"
	"github.com/gin-gonic/gin"
)

func QueryPatientByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Patient found",
		"id":      id,
	})
}

func AddPatient(c *gin.Context) {
	var input models.Patient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	username, _ := c.Get("username")
	hospital, _ := c.Get("hospital")

	c.JSON(http.StatusOK, gin.H{
		"message":  input,
		"user_id":  userID,
		"username": username,
		"hospital": hospital,
	})
}

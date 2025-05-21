package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryPatientByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Patient found",
		"id":      id,
	})
}

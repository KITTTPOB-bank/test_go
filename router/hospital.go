package router

import (
	"github.com/KITTTPOB-bank/hospitalapi/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	patientGroup := r.Group("/patient")
	{
		patientGroup.GET("/search/:id", controller.QueryPatientByID)
		// patientGroup.POST("/search")
	}

	staffGroup := r.Group("/staff")
	{
		staffGroup.POST("/create", controller.CreateStaff)
		staffGroup.POST("/login", controller.LoginStaff)
	}
}

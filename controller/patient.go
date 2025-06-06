package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KITTTPOB-bank/hospitalapi/models"
	"github.com/KITTTPOB-bank/hospitalapi/repository"
	"github.com/KITTTPOB-bank/hospitalapi/usecase"
	"github.com/gin-gonic/gin"
)

// QueryPatientByID godoc
// @Summary Query patient by ID
// @Description ดึงข้อมูลผู้ป่วยโดยใช้ national_id หรือ passport_id
// @Tags patient
// @Accept json
// @Produce json
// @Param id path string true "Passport ID OR National ID"
// @Router /patient/search/{id} [get]
func QueryPatientByID(c *gin.Context) {
	idVal := c.Param("id")

	patientRepo := repository.NewPatientRepository()
	userRepo := repository.NewUserRepository()
	patientUC := usecase.NewPatientUsecase(patientRepo, userRepo)
	patient, err := patientUC.SearchPatient(idVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"first_name_th":  patient.FirstNameTH,
		"middle_name_th": patient.MiddleNameTH,
		"last_name_th":   patient.LastNameTH,
		"first_name_en":  patient.FirstNameEN,
		"middle_name_en": patient.MiddleNameEN,
		"last_name_en":   patient.LastNameEN,
		"date_of_birth":  patient.DateOfBirth,
		"patient_hn":     patient.PatientHN,
		"national_id":    patient.NationalID,
		"passport_id":    patient.PassportID,
		"phone_number":   patient.PhoneNumber,
		"email":          patient.Email,
		"gender":         patient.Gender,
	})

}

// QueryPatientByStaff godoc
// @Summary Query patient by staff criteria
// @Description ค้นหาผู้ป่วยโดยอ้างอิงจาก รหัส staff
// @Tags patient
// @Accept json
// @Produce json
// @Param request body models.PatientRequest true "Patient search criteria"
// @Router /patient/search [post]
// @Security BearerAuth
func QueryPatientByStaff(c *gin.Context) {
	var input models.PatientRequest
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	usernameVal, _ := c.Get("username")
	hospitalVal, _ := c.Get("hospital")

	username, _ := usernameVal.(string)
	hospital, _ := hospitalVal.(string)

	patientRepo := repository.NewPatientRepository()
	userRepo := repository.NewUserRepository()
	patientUC := usecase.NewPatientUsecase(patientRepo, userRepo)

	msg, err := patientUC.SearchPatientByStaff(input, username, hospital)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": msg})

}

// AddPatient godoc
// @Summary Add multiple patients
// @Description เพิ่มผู้ป่วย
// @Tags staff
// @Accept json
// @Produce json
// @Param request body models.CreatePatientRequestGroup true "List of new patients"
// @Router /staff/patients [post]
// @Security BearerAuth
func AddPatient(c *gin.Context) {
	var input struct {
		Data []models.CreatePatientRequest `json:"data"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usernameVal, _ := c.Get("username")
	hospitalVal, _ := c.Get("hospital")
	username, _ := usernameVal.(string)
	hospital, _ := hospitalVal.(string)

	patientRepo := repository.NewPatientRepository()
	userRepo := repository.NewUserRepository()
	patientUC := usecase.NewPatientUsecase(patientRepo, userRepo)

	var results []string
	for i, p := range input.Data {
		msg, err := patientUC.AddNewPatient(p, username, hospital)
		if err != nil {
			results = append(results, fmt.Sprintf("index %d: error : %s", i, err))
			continue
		}
		results = append(results, fmt.Sprintf("index %d: %s", i, msg))
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

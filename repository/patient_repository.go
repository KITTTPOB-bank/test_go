package repository

import "github.com/KITTTPOB-bank/hospitalapi/models"

type PatientRepository interface {
	FindById(national_id, passport_id string) (*models.Patient, error)
	GeneratePatient(patient *models.Patient) error
	FindPatients(input models.PatientRequest, staff_id uint) ([]*models.Patient, error)
}

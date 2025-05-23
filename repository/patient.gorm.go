package repository

import (
	"github.com/KITTTPOB-bank/hospitalapi/initializers"
	"github.com/KITTTPOB-bank/hospitalapi/models"
)

type patientRepositoryGorm struct{}

func NewPatientRepository() PatientRepository {
	return &patientRepositoryGorm{}
}

func (r *patientRepositoryGorm) FindById(national_id, passport_id string) (*models.Patient, error) {
	var patient models.Patient
	result := initializers.DB.Where("national_id = ? OR passport_id = ?", national_id, passport_id).Limit(1).Find(&patient)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &patient, nil

}

func (r *patientRepositoryGorm) GeneratePatient(patient *models.Patient) error {
	return initializers.DB.Create(patient).Error
}
func (r *patientRepositoryGorm) FindPatients(input models.PatientRequest, staffID uint) ([]*models.Patient, error) {
	var patients []*models.Patient

	query := initializers.DB.Model(&models.Patient{}).Where("staff_id = ?", staffID)

	if input.FirstName != "" {
		query = query.Where("first_name_th LIKE ?", "%"+input.FirstName+"%")
	}
	if input.MiddleName != "" {
		query = query.Where("middle_name_th LIKE ?", "%"+input.MiddleName+"%")
	}
	if input.LastName != "" {
		query = query.Where("last_name_th LIKE ?", "%"+input.LastName+"%")
	}
	if input.DateOfBirth != "" {
		query = query.Where("date_of_birth = ?", input.DateOfBirth)
	}
	if input.NationalID != "" {
		query = query.Where("national_id = ?", input.NationalID)
	}
	if input.PassportID != "" {
		query = query.Where("passport_id = ?", input.PassportID)
	}
	if input.PhoneNumber != "" {
		query = query.Where("phone_number LIKE ?", "%"+input.PhoneNumber+"%")
	}
	if input.Email != "" {
		query = query.Where("email LIKE ?", "%"+input.Email+"%")
	}

	if err := query.Find(&patients).Error; err != nil {
		return nil, err
	}

	return patients, nil
}

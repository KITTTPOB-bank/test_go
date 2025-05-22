package usecase

import (
	"errors"
	"time"

	"github.com/KITTTPOB-bank/hospitalapi/initializers"
	"github.com/KITTTPOB-bank/hospitalapi/models"
	"github.com/KITTTPOB-bank/hospitalapi/repository"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

type PatientUsecase interface {
	AddNewPatient(req models.CreatePatientRequest, username, hospital string) (string, error)
	SearchPatient(id string) (*models.Patient, error)
	SearchPatientByStaff(input models.PatientRequest, username, hospital string) ([]*models.Patient, error)
}

type patientUsecase struct {
	repo     repository.PatientRepository
	userRepo repository.UserRepository
}

func NewPatientUsecase(pRepo repository.PatientRepository, uRepo repository.UserRepository) PatientUsecase {
	return &patientUsecase{repo: pRepo, userRepo: uRepo}
}

func (u *patientUsecase) AddNewPatient(req models.CreatePatientRequest, username, hospital string) (string, error) {
	user, err := u.userRepo.FindByUsername(username, hospital)
	if err != nil || user == nil {
		return "", errors.New("staff not found")
	}

	staffID := user.ID

	existing, _ := u.repo.FindById(req.NationalID, req.PassportID)
	if existing != nil {
		return "", errors.New("patient already exists")
	}

	now := time.Now()

	newPatient := &models.Patient{
		FirstNameTH:  req.FirstNameTH,
		MiddleNameTH: req.MiddleNameTH,
		LastNameTH:   req.LastNameTH,
		FirstNameEN:  req.FirstNameEN,
		MiddleNameEN: req.MiddleNameEN,
		LastNameEN:   req.LastNameEN,
		DateOfBirth:  req.DateOfBirth,
		PatientHN:    req.PatientHN,
		NationalID:   req.NationalID,
		PassportID:   req.PassportID,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		Gender:       req.Gender,
		StaffID:      staffID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := u.repo.GeneratePatient(newPatient); err != nil {
		return "", err
	}

	return "Created Sussed", nil
}

func (u *patientUsecase) SearchPatient(id string) (*models.Patient, error) {
	patient, err := u.repo.FindById(id, id)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (u *patientUsecase) SearchPatientByStaff(input models.PatientRequest, username, hospital string) ([]*models.Patient, error) {

	user, err := u.userRepo.FindByUsername(username, hospital)

	if err != nil || user == nil {
		return nil, errors.New("staff not found")
	}

	staffID := user.ID

	existing, err := u.repo.FindPatients(input, staffID)

	if err != nil {
		return nil, err
	}

	return existing, nil
}

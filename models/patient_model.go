package models

import "time"

type CreatePatientRequest struct {
	FirstNameTH  string `json:"first_name_th" binding:"required"`
	MiddleNameTH string `json:"middle_name_th,omitempty"`
	LastNameTH   string `json:"last_name_th" binding:"required"`
	FirstNameEN  string `json:"first_name_en,omitempty"`
	MiddleNameEN string `json:"middle_name_en,omitempty"`
	LastNameEN   string `json:"last_name_en,omitempty"`
	DateOfBirth  string `json:"date_of_birth,omitempty" binding:"omitempty,datetime=2006-01-02"`
	PatientHN    string `json:"patient_hn,omitempty"`
	NationalID   string `json:"national_id,omitempty"`
	PassportID   string `json:"passport_id,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	Email        string `json:"email,omitempty" binding:"omitempty,email"`
	Gender       string `json:"gender,omitempty" binding:"omitempty,oneof=M F"`
}

type PatientRequest struct {
	FirstName   string `json:"first_name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	DateOfBirth string `json:"date_of_birth,omitempty" binding:"omitempty,datetime=2006-01-02"`
	NationalID  string `json:"national_id,omitempty"`
	PassportID  string `json:"passport_id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty" binding:"omitempty,email"`
}

type Patient struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstNameTH  string    `json:"first_name_th" gorm:"type:text;not null"`
	MiddleNameTH string    `json:"middle_name_th" gorm:"type:text"`
	LastNameTH   string    `json:"last_name_th" gorm:"type:text;not null"`
	FirstNameEN  string    `json:"first_name_en" gorm:"type:text"`
	MiddleNameEN string    `json:"middle_name_en" gorm:"type:text"`
	LastNameEN   string    `json:"last_name_en" gorm:"type:text"`
	DateOfBirth  string    `json:"date_of_birth" gorm:"type:text"`
	PatientHN    string    `json:"patient_hn" gorm:"type:text"`
	NationalID   string    `json:"national_id" gorm:"type:varchar(13);unique"`
	PassportID   string    `json:"passport_id" gorm:"type:varchar(20);unique"`
	PhoneNumber  string    `json:"phone_number" gorm:"type:varchar(20)"`
	Email        string    `json:"email" gorm:"type:text"`
	Gender       string    `json:"gender" gorm:"type:char(1);check:gender IN ('M','F')"`
	StaffID      uint      `json:"staff_id"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Patient) TableName() string {
	return "patient"
}

package models

type Patient struct {
	FirstNameTH  string `json:"first_name_th" binding:"required"`
	MiddleNameTH string `json:"middle_name_th,omitempty"`
	LastNameTH   string `json:"last_name_th" binding:"required"`
	FirstNameEN  string `json:"first_name_en,omitempty"`
	MiddleNameEN string `json:"middle_name_en,omitempty"`
	LastNameEN   string `json:"last_name_en,omitempty"`
	DateOfBirth  string `json:"date_of_birth" binding:"required,datetime=2006-01-02"`
	PatientHN    string `json:"patient_hn,omitempty"`
	NationalID   string `json:"national_id,omitempty"`
	PassportID   string `json:"passport_id,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	Email        string `json:"email,omitempty" binding:"omitempty,email"`
	Gender       string `json:"gender" binding:"required,oneof=M F"`
}

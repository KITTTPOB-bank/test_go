package test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/KITTTPOB-bank/hospitalapi/models"
	"github.com/KITTTPOB-bank/hospitalapi/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type MockPatientRepository struct{}

func (m *MockPatientRepository) FindById(nationalID, passportID string) (*models.Patient, error) {
	if nationalID == "1234567890123" {
		return &models.Patient{
			FirstNameTH: "สมชาย",
			PatientHN:   "HN001",
			NationalID:  "1234567890123",
		}, nil
	}
	return nil, nil
}
func (m *MockPatientRepository) GeneratePatient(patient *models.Patient) error {

	return nil
}

type MockUserRepository struct{}

func (m *MockUserRepository) GenerateStaff(user *models.User) error {
	return nil
}

func (m *MockUserRepository) FindByUsername(username, hospital string) (*models.User, error) {
	if username == "testuser" && hospital == "testhospital" {
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("pass0101"), bcrypt.DefaultCost)
		return &models.User{ID: 1, Username: username, Password: string(passwordHash), Hospital: hospital}, nil
	}
	return nil, errors.New("user not found")

}

func (m *MockPatientRepository) FindPatients(input models.PatientRequest, staffID uint) ([]*models.Patient, error) {
	patients := []*models.Patient{
		{
			ID:           13,
			FirstNameTH:  "ชุติมา",
			MiddleNameTH: "รัตนาพร",
			LastNameTH:   "แก้วใส",
			FirstNameEN:  "Chutima",
			MiddleNameEN: "Rattanaporn",
			LastNameEN:   "Kaewsai",
			DateOfBirth:  "1995-06-18T00:00:00Z",
			PatientHN:    "HN334455",
			NationalID:   "1400508765432",
			PassportID:   "P019283746",
			PhoneNumber:  "0912345670",
			Email:        "chutima@example.com",
			Gender:       "F",
			StaffID:      staffID,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}
	if staffID == 1 && input.Email == "chutima@example.com" {
		return patients, nil
	}
	return nil, nil

}

// "/patient/search"
func setupRouterQueryPatientByStaff() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	pRepo := &MockPatientRepository{}
	uRepo := &MockUserRepository{}
	patientUC := usecase.NewPatientUsecase(pRepo, uRepo)

	r.POST("/patient/search", mockAuthorization(), func(c *gin.Context) {
		var input models.PatientRequest

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			return
		}
		username, _ := c.Get("username")
		hospital, _ := c.Get("hospital")

		msg, err := patientUC.SearchPatientByStaff(input, username.(string), hospital.(string))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": msg})
	})

	return r
}

func TestQueryPatientByStaff_Success(t *testing.T) {
	router := setupRouterQueryPatientByStaff()

	jsonStr := `{
		"email": "chutima@example.com"
	}`
	req, _ := http.NewRequest("POST", "/patient/search", strings.NewReader(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", "Bearer "+generateTestToken())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "message")
}

func TestQueryPatientByStaff_NotSuccess(t *testing.T) {
	router := setupRouterQueryPatientByStaff()

	req, _ := http.NewRequest("POST", "/patient/search", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+generateTestToken())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid request body")
}

// "/patient/search/:id"
func setupRouterQueryPatientByID() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	pRepo := &MockPatientRepository{}
	uRepo := &MockUserRepository{}
	patientUC := usecase.NewPatientUsecase(pRepo, uRepo)

	r.GET("/patient/search/:id", func(c *gin.Context) {
		id := c.Param("id")
		patient, err := patientUC.SearchPatient(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if patient == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"first_name_th": patient.FirstNameTH,
			"patient_hn":    patient.PatientHN,
			"national_id":   patient.NationalID,
		})
	})

	return r
}

func mockAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("username", "testuser")
		c.Set("hospital", "testhospital")
		c.Next()
	}
}
func generateTestToken() string {
	secret := []byte("MAPLE_SYRUB")

	claims := jwt.MapClaims{
		"username": "testuser",
		"hospital": "testhospital",
		"exp":      time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(secret)
	return tokenString
}

func TestQueryPatientByID_Success(t *testing.T) {
	router := setupRouterQueryPatientByID()

	req, _ := http.NewRequest("GET", "/patient/search/1234567890123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "สมชาย")
	assert.Contains(t, w.Body.String(), "HN001")
}

func TestQueryPatientByID_NotFound(t *testing.T) {
	router := setupRouterQueryPatientByID()

	req, _ := http.NewRequest("GET", "/patient/search/0000000000000", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "patient not found")
}

// "/staff/create"
func setupRouterCreateStaff() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	uRepo := &MockUserRepository{}
	userUC := usecase.NewUserUsecase(uRepo)

	r.POST("/staff/create", func(c *gin.Context) {
		var staff models.Authentication

		if err := c.ShouldBindJSON(&staff); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		msg, err := userUC.Register(staff.Username, staff.Password, staff.Hospital)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": msg,
		})
	})

	return r
}

func TestCreateStaff_Success(t *testing.T) {
	router := setupRouterCreateStaff()

	jsonStr := `{
		"username": "cc",
		"password": "pass0101",
		"hospital": "testhospital"
	}`
	req, _ := http.NewRequest("POST", "/staff/create", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}

func TestCreateStaff_alreadyExists(t *testing.T) {
	router := setupRouterCreateStaff()

	jsonStr := `{
		"username": "testuser",
		"password": "pass0101",
		"hospital": "testhospital"
	}`
	req, _ := http.NewRequest("POST", "/staff/create", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "username already used")
}

// "/staff/login"
func setupRouterLoginStaff() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	uRepo := &MockUserRepository{}
	userUC := usecase.NewUserUsecase(uRepo)

	r.POST("/staff/login", func(c *gin.Context) {
		var staff models.Authentication

		if err := c.ShouldBindJSON(&staff); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		msg, err := userUC.Login(staff.Username, staff.Password, staff.Hospital)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": msg,
		})
	})

	return r
}

func TestLogin_Success(t *testing.T) {
	router := setupRouterLoginStaff()

	jsonStr := `{
		"username": "testuser",
		"password": "pass0101",
		"hospital": "testhospital"
	}`
	req, _ := http.NewRequest("POST", "/staff/login", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}

func TestLogin_WrongPassword(t *testing.T) {
	router := setupRouterLoginStaff()

	jsonStr := `{
		"username": "testuser",
		"password": "xxxxx",
		"hospital": "testhospital"
	}`
	req, _ := http.NewRequest("POST", "/staff/login", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid Password")
}

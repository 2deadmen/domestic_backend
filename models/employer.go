package models

import (
	"log"
	"time"

	"github.com/2deadmen/domestic_backend/services"
	"github.com/pkg/errors"
)

//  models

// real
type Employer struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Email        string    `json:"email" gorm:"unique"`
	Password     string    `json:"password"`
	Age          int       `json:"age"`
	Gender       string    `json:"gender"`
	Phone        string    `json:"phone" `
	AddressProof string    `json:"addressproof"`
	Type         string    `json:"type"`
	OTP          string    `json:"otp"`
	Verified     bool      `json:"verified" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type Employee struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Pin            string    `json:"pin" `
	Age            int       `json:"age"`
	Gender         string    `json:"gender"`
	Phone          string    `json:"phone" `
	AddressProof   string    `json:"addressproof"`
	OpenToWork     bool      `json:"opentowork"`
	WorkExperience string    `json:"workexperience"`
	TypeOfWork     []string  `json:"typeofwork" gorm:"type:json"`
	PhotoURL       string    `json:"photourl"`
	Verified       bool      `json:"verified" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// MigrateModels ensures that the database schema matches the models
func MigrateModels() {
	log.Println("Starting database migrations...")

	// Perform auto-migration for all models
	err := services.DB.AutoMigrate(&Employee{}, &Employer{})
	if err != nil {
		log.Fatalf("Failed to migrate database models: %v", err)
	}

	log.Println("Database migrations completed successfully.")
}

// CheckEmailExists checks if an email is already registered in the database
func CheckEmailExists(email string) (bool, error) {
	var count int64
	err := services.DB.Model(&Employer{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CreateEmployer inserts a new employer into the database
func CreateEmployer(employer *Employer) error {
	// Insert the employer and automatically populate the ID
	if err := services.DB.Create(employer).Error; err != nil {
		return errors.Wrap(err, "failed to create employer")
	}
	return nil
}

// get employer by mail
func GetEmployerByEmail(email string) (Employer, error) {
	var employer Employer
	err := services.DB.Where("email = ?", email).First(&employer).Error
	return employer, err
}

// update employer
func UpdateEmployer(employer *Employer) error {
	return services.DB.Save(employer).Error
}

// GetAllEmployers retrieves all employers from the database
func GetAllEmployers() ([]Employer, error) {
	var employers []Employer
	err := services.DB.Find(&employers).Error
	if err != nil {
		return nil, err
	}
	return employers, nil
}

// GetEmployerByID retrieves an employer by ID
func GetEmployerByID(id string) (Employer, error) {
	var employer Employer
	err := services.DB.First(&employer, "id = ?", id).Error
	return employer, err
}

// UpdateEmployerByID updates an employer's details by ID
func UpdateEmployerByID(id string, updatedEmployer *Employer) error {
	var employer Employer
	if err := services.DB.First(&employer, "id = ?", id).Error; err != nil {
		return err
	}

	return services.DB.Model(&employer).Updates(updatedEmployer).Error
}

// DeleteEmployerByID deletes an employer by ID
func DeleteEmployerByID(id string) error {
	return services.DB.Delete(&Employer{}, "id = ?", id).Error
}

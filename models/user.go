package models

import (
	"log"
	"time"

	"github.com/2deadmen/domestic_backend/services"
	"github.com/pkg/errors"
)

//  models

// dummy
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

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
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type Employee struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Pin            uint64    `json:"pin" `
	Age            int       `json:"age"`
	Gender         string    `json:"gender"`
	Phone          string    `json:"phone" `
	AddressProof   string    `json:"addressproof"`
	OpenToWork     bool      `json:"opentowork"`
	WorkExperience string    `json:"workexperience"`
	TypeOfWork     []string  `json:"typeofwork" gorm:"type:json"`
	PhotoURL       string    `json:"photourl"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// MigrateModels ensures that the database schema matches the models
func MigrateModels() {
	log.Println("Starting database migrations...")

	// Perform auto-migration for all models
	err := services.DB.AutoMigrate(&User{}, &Employee{}, &Employer{})
	if err != nil {
		log.Fatalf("Failed to migrate database models: %v", err)
	}

	log.Println("Database migrations completed successfully.")
}

// Get all users
func GetAllUsers() ([]User, error) {
	var users []User
	err := services.DB.Find(&users).Error
	return users, err
}

// Create a new user
func CreateUser(user *User) error {
	return services.DB.Create(user).Error
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

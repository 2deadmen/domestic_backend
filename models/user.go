package models

import (
	"log"
	"time"

	"github.com/2deadmen/domestic_backend/services"
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

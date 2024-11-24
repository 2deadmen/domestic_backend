package models

import (
	"log"

	"github.com/2deadmen/domestic_backend/services"
)

// User model
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

// MigrateModels ensures that the database schema matches the models
func MigrateModels() {
	log.Println("Starting database migrations...")

	// Perform auto-migration for all models
	err := services.DB.AutoMigrate(&User{})
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

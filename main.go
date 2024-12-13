package main

import (
	"log"

	"github.com/2deadmen/domestic_backend/config"
	"github.com/2deadmen/domestic_backend/routes"
	"github.com/2deadmen/domestic_backend/services"
)

func main() {
	// Initialize configurations
	config.LoadConfig()

	// Initialize database connection
	services.InitDB()

	// Check if DB connection is established
	if services.DB == nil {
		log.Fatal("Database connection failed!")
	}

	// Initialize router with routes
	router := routes.InitRoutes()

	// CORS middleware configuration

	// Start the server
	log.Println("Server running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

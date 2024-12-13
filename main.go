package main

import (
	"log"

	"github.com/2deadmen/domestic_backend/config"
	"github.com/2deadmen/domestic_backend/routes"
	"github.com/2deadmen/domestic_backend/services"
	"github.com/gin-contrib/cors"
)

func main() {
	// Initialize configurations
	config.LoadConfig()

	services.InitDB()

	// Check if DB connection is established
	if services.DB == nil {
		log.Fatal("Database connection failed!")
	}
	// Initialize router with routes
	router := routes.InitRoutes()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},                   // Explicitly allow your frontend origin
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow all methods
		AllowHeaders: []string{"*"},                                       // Allow all headers
	}))

	// Start the server
	log.Println("Server running on port 8080")
	router.Run(":8080")
}

package controllers

import (
	"net/http"

	"github.com/2deadmen/domestic_backend/models"
	"github.com/2deadmen/domestic_backend/utils"

	"github.com/gin-gonic/gin"
)

// Get all users
func GetUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
		return
	}
	utils.RespondJSON(c, http.StatusOK, users)
}

// Create a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	models.MigrateModels()

	if err := models.CreateUser(&user); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
		return
	}

	utils.RespondJSON(c, http.StatusCreated, user)
}

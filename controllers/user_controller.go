package controllers

import (
	"net/http"

	"github.com/2deadmen/domestic_backend/models"
	"github.com/2deadmen/domestic_backend/utils"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get all users
// @Description Fetch all users from the database
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func GetUsers(c *gin.Context) {

	// Fetch all users
	users, err := models.GetAllUsers()
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
		return
	}

	// Respond with the combined data
	utils.RespondJSON(c, http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Add a new user to the database
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := models.CreateUser(&user); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
		return
	}

	utils.RespondJSON(c, http.StatusCreated, user)
}

// CreateEmployer godoc
// @Summary Register a new employer
// @Description Sign up a new employer with email and password
// @Tags Employers
// @Accept json
// @Produce json
// @Param employer body models.Employer true "Employer Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employers [post]
func CreateEmployer(c *gin.Context) {
	var employer models.Employer

	// Bind JSON input to the Employer model
	if err := c.ShouldBindJSON(&employer); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Validate required fields
	if employer.Email == "" || employer.Password == "" {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Check if email already exists
	if exists, err := models.CheckEmailExists(employer.Email); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to validate email"})
		return
	} else if exists {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Email is already registered"})
		return
	}

	// Hash the password for security
	hashedPassword, err := utils.HashPassword(employer.Password)
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}
	employer.Password = hashedPassword

	// Save the employer to the database
	if err := models.CreateEmployer(&employer); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Unable to create employer"})
		return
	}

	// Generate a JWT for the newly created employer
	token, err := utils.GenerateJWT(employer.ID, employer.Email)
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with a success message, employer email, and JWT
	utils.RespondJSON(c, http.StatusCreated, gin.H{
		"message": "Employer created successfully",
		"id":      employer.ID,
		"token":   token,
	})
}

// GetAllEmployers godoc
// @Summary Get all employers
// @Description Retrieve a list of all employers
// @Tags Employers
// @Produce json
// @Success 200 {array} models.Employer
// @Failure 500 {object} map[string]string
// @Router /employers [get]
func GetAllEmployers(c *gin.Context) {
	employers, err := models.GetAllEmployers()
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to fetch employers"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, employers)
}

// GetEmployer godoc
// @Summary Get an employer by ID
// @Description Retrieve an employer's details
// @Tags Employers
// @Produce json
// @Param id path int true "Employer ID"
// @Success 200 {object} models.Employer
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employers/{id} [get]
func GetEmployer(c *gin.Context) {
	id := c.Param("id")

	employer, err := models.GetEmployerByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Employer not found"})
			return
		}
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to fetch employer"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, employer)
}

// UpdateEmployer godoc
// @Summary Update an employer
// @Description Modify an employer's details
// @Tags Employers
// @Accept json
// @Produce json
// @Param id path int true "Employer ID"
// @Param employer body models.Employer true "Updated Employer Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employers/{id} [put]
func UpdateEmployer(c *gin.Context) {
	id := c.Param("id")
	var updatedEmployer models.Employer

	if err := c.ShouldBindJSON(&updatedEmployer); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := models.UpdateEmployerByID(id, &updatedEmployer); err != nil {
		if err.Error() == "record not found" {
			utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Employer not found"})
			return
		}
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to update employer"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Employer updated successfully"})
}

// DeleteEmployer godoc
// @Summary Delete an employer
// @Description Remove an employer by ID
// @Tags Employers
// @Produce json
// @Param id path int true "Employer ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employers/{id} [delete]
func DeleteEmployer(c *gin.Context) {
	id := c.Param("id")

	if err := models.DeleteEmployerByID(id); err != nil {
		if err.Error() == "record not found" {
			utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Employer not found"})
			return
		}
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to delete employer"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Employer deleted successfully"})
}

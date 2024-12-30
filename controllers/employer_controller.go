package controllers

import (
	"fmt"
	"net/http"

	"github.com/2deadmen/domestic_backend/models"
	"github.com/2deadmen/domestic_backend/utils"

	"github.com/gin-gonic/gin"
)

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

	// Generate an OTP and set it in the employer record
	otp := utils.GenerateOTP() // Implement this utility function
	employer.OTP = otp

	// Send OTP to the employer's email
	if err := utils.SendEmail(employer.Email, "Your OTP", fmt.Sprintf("Your OTP is: %s", otp)); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to send OTP email"})
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
	token, err := utils.GenerateJWT(employer.ID, employer.Email)
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	utils.RespondJSON(c, http.StatusCreated, gin.H{
		"message": "Employer created successfully. OTP has been sent to your email.",
		"id":      employer.ID,
		"token":   token,
	})

}

// VerifyOTP godoc
// @Summary Verify an employer's OTP
// @Description Verify OTP for an employer and mark as verified
// @Tags Employers
// @Accept json
// @Produce json
// @Param otp body map[string]string true "OTP and Email"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employers/verify-otp [post]
func VerifyOTP(c *gin.Context) {
	var requestBody struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	// Bind JSON input to the request body struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Find the employer by email
	employer, err := models.GetEmployerByEmail(requestBody.Email)
	if err != nil {
		utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Employer not found"})
		return
	}

	// Check if the OTP matches
	if employer.OTP != requestBody.OTP {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
		return
	}

	// Update the Verified field to true
	employer.Verified = true
	employer.OTP = "" // Clear the OTP after successful verification
	if err := models.UpdateEmployer(&employer); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to update employer status"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(employer.ID, employer.Email)
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with success message and token
	utils.RespondJSON(c, http.StatusOK, gin.H{
		"message": "OTP verified successfully",
		"id":      employer.ID,
		"token":   token,
	})
}

// SignIn godoc
// @Summary Sign in an employer
// @Description Log in an employer with email and password
// @Tags Employers
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Email and Password"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employers/sign-in [post]
func SignIn(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON input to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Find the employer by email
	employer, err := models.GetEmployerByEmail(credentials.Email)
	if err != nil {
		utils.RespondJSON(c, http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check if the employer is verified
	if !employer.Verified {
		utils.RespondJSON(c, http.StatusUnauthorized, gin.H{"error": "Email not verified"})
		return
	}

	// Verify the password
	if !utils.CheckPasswordHash(credentials.Password, employer.Password) {
		utils.RespondJSON(c, http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(employer.ID, employer.Email)
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with success message and token
	utils.RespondJSON(c, http.StatusOK, gin.H{
		"message": "Sign-in successful",
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

package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/2deadmen/domestic_backend/models"
	"github.com/2deadmen/domestic_backend/utils"
	"github.com/gin-gonic/gin"
)

// @Summary update job card status rejected/accepted
func HandleJobApplicationStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}

	var requestBody struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if requestBody.Status != "accepted" && requestBody.Status != "rejected" {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Status must be 'accepted' or 'rejected'"})
		return

	}

	if err := models.UpdateJobApplicationStatusByID(id, requestBody.Status); err != nil {
		if err.Error() == "record not found" {
			utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Job application not found"})
			return
		}
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to update application status"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Application status updated successfully"})
}

// CreateJobApplicationHandler handles creating a new job application
func CreateJobApplicationHandler(c *gin.Context) {
	var jobApplication models.JobApplication

	// Bind the request body to the job application model
	if err := c.ShouldBindJSON(&jobApplication); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Fetch employer details from the Employer table
	employer, err := models.GetEmployerByID(strconv.Itoa(jobApplication.EmployerId))
	if err != nil {
		utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Employer not found"})
		return
	}

	// Fetch employee details from the Employee table
	employee, err := models.GetEmployeeByID(strconv.Itoa(jobApplication.EmployeeId))
	if err != nil {
		utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	// Send email to the employer with job application details
	emailContent := fmt.Sprintf(
		"New Job Application:\n\nEmployee Name: %s: %s%s%s%s\nEmployee Phone: %s\n\nPlease review the application.",
		employee.Name, employee.Gender, employee.TypeOfWork, employee.Phone, employee.CreatedAt, employee.WorkExperience,
	)

	if err := utils.SendEmail(employer.Email, "New Job Application", emailContent); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to send email to employer"})
		return
	}

	// Create the job application record in the database
	if err := models.CreateJobApplication(&jobApplication); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to create job application"})
		return
	}

	utils.RespondJSON(c, http.StatusCreated, gin.H{"message": "Job application created successfully"})
}

// DeleteJobApplicationHandler handles deleting a job application by ID
func DeleteJobApplicationHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid job application ID"})
		return
	}

	if err := models.DeleteJobApplication(id); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to delete job application"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Job application deleted successfully"})
}

func GetApplicationsByEmployerOrEmployeeHandler(c *gin.Context) {
	userType := c.Query("userType")
	userId := c.Query("userId")

	var applications []models.JobApplication
	var err error

	if userType == "Employer" {
		applications, err = models.GetApplicationsByEmployer(userId)
	} else if userType == "Employee" {
		applications, err = models.GetApplicationsByEmployee(userId)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching applications"})
		return
	}

	c.JSON(http.StatusOK, applications)
}

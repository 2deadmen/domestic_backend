package controllers

import (
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

	if err := c.ShouldBindJSON(&jobApplication); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

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

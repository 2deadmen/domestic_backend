package controllers

import (
	"net/http"
	"strconv"

	"github.com/2deadmen/domestic_backend/models"
	"github.com/2deadmen/domestic_backend/utils"
	"github.com/gin-gonic/gin"
)

// CreateJobCard creates a new job card
// @Summary Create a new job card
// @Description Creates a new job card based on the provided JSON data
// @Tags JobCard
// @Accept json
// @Produce json
// @Param jobCard body models.JobCard true "Job Card Data"
// @Success 201 {object} gin.H{"message": "Job card created successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid request body"}
// @Failure 500 {object} gin.H{"error": "Failed to create job card"}
// @Router /jobcards [post]
func CreateJobCard(c *gin.Context) {
	var jobCard models.JobCard
	if err := c.ShouldBindJSON(&jobCard); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := models.CreateJobCard(&jobCard); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to create job card"})
		return
	}

	utils.RespondJSON(c, http.StatusCreated, gin.H{"message": "Job card created successfully"})
}

// GetAllJobCards retrieves all job cards
// @Summary Retrieve all job cards
// @Description Fetches all job cards from the database
// @Tags JobCard
// @Accept json
// @Produce json
// @Success 200 {array} models.JobCard
// @Failure 500 {object} gin.H{"error": "Failed to fetch job cards"}
// @Router /jobcards [get]
func GetAllJobCards(c *gin.Context) {
	jobCards, err := models.GetAllJobCards()
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to fetch job cards"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, jobCards)
}

// GetJobCard retrieves a job card by ID
// @Summary Retrieve a job card by ID
// @Description Fetches a specific job card by its ID
// @Tags JobCard
// @Accept json
// @Produce json
// @Param id path int true "Job Card ID"
// @Success 200 {object} models.JobCard
// @Failure 400 {object} gin.H{"error": "Invalid job card ID"}
// @Failure 404 {object} gin.H{"error": "Job card not found"}
// @Failure 500 {object} gin.H{"error": "Failed to fetch job card"}
// @Router /jobcards/{id} [get]
func GetJobCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid job card ID"})
		return
	}

	jobCard, err := models.GetJobCardByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Job card not found"})
			return
		}
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to fetch job card"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, jobCard)
}

// UpdateJobCard updates a job card
// @Summary Update a job card by ID
// @Description Updates a job card using the provided JSON data and the specified ID
// @Tags JobCard
// @Accept json
// @Produce json
// @Param id path int true "Job Card ID"
// @Param jobCard body models.JobCard true "Updated Job Card Data"
// @Success 200 {object} gin.H{"message": "Job card updated successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid job card ID"}
// @Failure 404 {object} gin.H{"error": "Job card not found"}
// @Failure 500 {object} gin.H{"error": "Failed to update job card"}
// @Router /jobcards/{id} [put]
func UpdateJobCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid job card ID"})
		return
	}

	var updatedJobCard models.JobCard
	if err := c.ShouldBindJSON(&updatedJobCard); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := models.UpdateJobCardByID(id, &updatedJobCard); err != nil {
		if err.Error() == "record not found" {
			utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Job card not found"})
			return
		}
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to update job card"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Job card updated successfully"})
}

// DeleteJobCard deletes a job card by ID
// @Summary Delete a job card by ID
// @Description Deletes a specific job card by its ID
// @Tags JobCard
// @Accept json
// @Produce json
// @Param id path int true "Job Card ID"
// @Success 200 {object} gin.H{"message": "Job card deleted successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid job card ID"}
// @Failure 404 {object} gin.H{"error": "Job card not found"}
// @Failure 500 {object} gin.H{"error": "Failed to delete job card"}
// @Router /jobcards/{id} [delete]
func DeleteJobCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid job card ID"})
		return
	}

	if err := models.DeleteJobCardByID(id); err != nil {
		if err.Error() == "record not found" {
			utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Job card not found"})
			return
		}
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to delete job card"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Job card deleted successfully"})
}

// GetActiveJobCards retrieves all active job cards
// @Summary Retrieve all active job cards
// @Description Fetches all job cards that are currently active
// @Tags JobCard
// @Accept json
// @Produce json
// @Success 200 {array} models.JobCard
// @Failure 500 {object} gin.H{"error": "Failed to fetch active job cards"}
// @Router /jobcards/active [get]
func GetActiveJobCards(c *gin.Context) {
	jobCards, err := models.GetActiveJobCards()
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to fetch active job cards"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, jobCards)
}

// UpdateJobCardActiveStatus updates the active status of a JobCard
// @Summary Update the active status of a job card by ID
// @Description Updates the active status for the specified job card ID based on provided JSON data
// @Tags JobCard
// @Accept json
// @Produce json
// @Param id path int true "Job Card ID"
// @Param requestBody body struct { Active bool `json:"active"` } true "Active Status"
// @Success 200 {object} gin.H{"message": "Job card status updated successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid job card ID"}
// @Failure 404 {object} gin.H{"error": "Job card not found"}
// @Failure 500 {object} gin.H{"error": "Failed to update job card status"}
// @Router /jobcards/{id}/status [put]
func UpdateJobCardActiveStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid job card ID"})
		return
	}

	var requestBody struct {
		Active bool `json:"active"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := models.UpdateJobCardActiveStatusByID(id, requestBody.Active); err != nil {
		if err.Error() == "record not found" {
			utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Job card not found"})
			return
		}
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to update job card status"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Job card status updated successfully"})
}

// HandleJobApplicationStatus updates the status of a JobApplication (accepted/rejected)
// @Summary Update the status of a job application by ID
// @Description Updates the status of the job application (accepted/rejected) by its ID
// @Tags JobApplication
// @Accept json
// @Produce json
// @Param id path int true "Application ID"
// @Param requestBody body struct { Status string `json:"status"` } true "Application Status"
// @Success 200 {object} gin.H{"message": "Application status updated successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid application ID"}
// @Failure 400 {object} gin.H{"error": "Status must be 'accepted' or 'rejected'"}
// @Failure 404 {object} gin.H{"error": "Job application not found"}
// @Failure 500 {object} gin.H{"error": "Failed to update application status"}
// @Router /jobapplications/{id}/status [put]
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

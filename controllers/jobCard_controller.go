package controllers

import (
	"net/http"
	"strconv"

	"github.com/2deadmen/domestic_backend/models"
	"github.com/2deadmen/domestic_backend/utils"
	"github.com/gin-gonic/gin"
)

// @Summary create a job card
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

// @Summary get all  job cards
func GetAllJobCards(c *gin.Context) {
	jobCards, err := models.GetAllJobCards()
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to fetch job cards"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, jobCards)
}

// @Summary get a job card
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

// @Summary get a job card with a id
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

// @Summary Delete a job card
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

// @Summary get active  job cards
func GetActiveJobCards(c *gin.Context) {
	jobCards, err := models.GetActiveJobCards()
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to fetch active job cards"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, jobCards)
}

// @Summary update  job card active status
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

package controllers

import (
	"net/http"
	"strconv"

	"github.com/2deadmen/domestic_backend/models"
	"github.com/gin-gonic/gin"
)

// CreateJobCampaignController handles creating a new job campaign.
func CreateJobCampaignController(c *gin.Context) {
	var campaign models.JobCampaign

	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateJobCampaign(&campaign); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, campaign)
}

// GetAllJobCampaignsController retrieves all job campaigns.
func GetAllJobCampaignsController(c *gin.Context) {
	campaigns, err := models.GetAllJobCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}

// GetJobCampaignByIDController retrieves a job campaign by ID.
func GetJobCampaignByIDController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job campaign ID"})
		return
	}

	campaign, err := models.GetJobCampaignByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaign)
}

// UpdateJobCampaignController updates a job campaign by ID.
func UpdateJobCampaignController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job campaign ID"})
		return
	}

	var updatedCampaign models.JobCampaign
	if err := c.ShouldBindJSON(&updatedCampaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateJobCampaign(uint(id), &updatedCampaign); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "job campaign updated successfully"})
}

// DeleteJobCampaignController deletes a job campaign by ID.
func DeleteJobCampaignController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job campaign ID"})
		return
	}

	if err := models.DeleteJobCampaign(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "job campaign deleted successfully"})
}

// GetActiveJobCampaignsController retrieves all active job campaigns.
func GetActiveJobCampaignsController(c *gin.Context) {
	campaigns, err := models.GetActiveJobCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}

// CreateCampaignApplicationController handles creating a new campaign application.
func CreateCampaignApplicationController(c *gin.Context) {
	var application models.CampaignApplication

	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateCampaignApplication(&application); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, application)
}

// GetApplicationsByCampaignIDController retrieves applications for a specific campaign.
func GetApplicationsByCampaignIDController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("campaign_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid campaign ID"})
		return
	}

	applications, err := models.GetApplicationsByCampaignID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, applications)
}

// UpdateApplicationStatusController updates the status of a campaign application.
func UpdateApplicationStatusController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid application ID"})
		return
	}

	var requestBody struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateApplicationStatus(uint(id), requestBody.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "application status updated successfully"})
}

// DeleteCampaignApplicationController deletes a campaign application by ID.
func DeleteCampaignApplicationController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid application ID"})
		return
	}

	if err := models.DeleteCampaignApplication(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "campaign application deleted successfully"})
}

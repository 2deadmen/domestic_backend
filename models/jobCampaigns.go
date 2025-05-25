package models

import (
	"github.com/2deadmen/domestic_backend/services"
	"github.com/pkg/errors"
	"time"
)

// CreateJobCampaign inserts a new JobCampaign into the database.
func CreateJobCampaign(campaign *JobCampaign) error {
	if err := services.DB.Create(campaign).Error; err != nil {
		return errors.Wrap(err, "failed to create job campaign")
	}
	return nil
}

// GetAllJobCampaigns retrieves all JobCampaigns from the database.
func GetAllJobCampaigns() ([]JobCampaign, error) {
	var campaigns []JobCampaign
	if err := services.DB.Find(&campaigns).Error; err != nil {
		return nil, errors.Wrap(err, "failed to fetch job campaigns")
	}
	return campaigns, nil
}

// GetJobCampaignByID retrieves a specific JobCampaign by its ID.
func GetJobCampaignByID(id uint) (*JobCampaign, error) {
	var campaign JobCampaign
	if err := services.DB.First(&campaign, "id = ?", id).Error; err != nil {
		return nil, errors.Wrap(err, "failed to fetch job campaign")
	}
	return &campaign, nil
}

// UpdateJobCampaign updates the details of an existing JobCampaign.
func UpdateJobCampaign(id uint, updatedCampaign *JobCampaign) error {
	var campaign JobCampaign
	if err := services.DB.First(&campaign, "id = ?", id).Error; err != nil {
		return errors.Wrap(err, "job campaign not found")
	}

	if err := services.DB.Model(&campaign).Updates(updatedCampaign).Error; err != nil {
		return errors.Wrap(err, "failed to update job campaign")
	}
	return nil
}

// DeleteJobCampaign deletes a JobCampaign by its ID.
func DeleteJobCampaign(id uint) error {
	if err := services.DB.Delete(&JobCampaign{}, "id = ?", id).Error; err != nil {
		return errors.Wrap(err, "failed to delete job campaign")
	}
	return nil
}

// GetActiveJobCampaigns retrieves all active JobCampaigns.
func GetActiveJobCampaigns() ([]JobCampaign, error) {
	var campaigns []JobCampaign
	if err := services.DB.Where("active = ?", true).Find(&campaigns).Error; err != nil {
		return nil, errors.Wrap(err, "failed to fetch active job campaigns")
	}
	return campaigns, nil
}

// CloseExpiredCampaigns deactivates campaigns that have passed their end date.
func CloseExpiredCampaigns() error {
	if err := services.DB.Model(&JobCampaign{}).
		Where("end_date < ?", time.Now()).
		Update("active", false).Error; err != nil {
		return errors.Wrap(err, "failed to close expired campaigns")
	}
	return nil
}

// CreateCampaignApplication adds a new CampaignApplication record.
func CreateCampaignApplication(application *CampaignApplication) error {
	if err := services.DB.Create(application).Error; err != nil {
		return errors.Wrap(err, "failed to create campaign application")
	}
	return nil
}

// GetApplicationsByCampaignID retrieves all applications for a specific JobCampaign.
func GetApplicationsByCampaignID(jobCampaignID uint) ([]CampaignApplication, error) {
	var applications []CampaignApplication
	if err := services.DB.Where("job_campaign_id = ?", jobCampaignID).Find(&applications).Error; err != nil {
		return nil, errors.Wrap(err, "failed to fetch campaign applications")
	}
	return applications, nil
}

// GetApplicationByID retrieves a specific CampaignApplication by its ID.
func GetApplicationByID(id uint) (*CampaignApplication, error) {
	var application CampaignApplication
	if err := services.DB.First(&application, "id = ?", id).Error; err != nil {
		return nil, errors.Wrap(err, "failed to fetch campaign application")
	}
	return &application, nil
}

// UpdateApplicationStatus updates the status of a CampaignApplication.
func UpdateApplicationStatus(id uint, status string) error {
	var application CampaignApplication
	if err := services.DB.First(&application, "id = ?", id).Error; err != nil {
		return errors.Wrap(err, "campaign application not found")
	}

	if err := services.DB.Model(&application).Update("status", status).Error; err != nil {
		return errors.Wrap(err, "failed to update application status")
	}
	return nil
}

// DeleteCampaignApplication deletes a CampaignApplication by its ID.
func DeleteCampaignApplication(id uint) error {
	if err := services.DB.Delete(&CampaignApplication{}, "id = ?", id).Error; err != nil {
		return errors.Wrap(err, "failed to delete campaign application")
	}
	return nil
}

package models

import (
	"github.com/2deadmen/domestic_backend/services"
	"github.com/pkg/errors"
)

// UpdateJobApplicationStatusByID updates the status of a JobApplication.
func UpdateJobApplicationStatusByID(id int, status string) error {
	if status != "accepted" && status != "rejected" {
		return errors.New("invalid status value")
	}

	if err := services.DB.Model(&JobApplication{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return errors.Wrap(err, "failed to update job application status")
	}
	return nil
}

// CreateJobApplication adds a new job application to the database
func CreateJobApplication(jobApplication *JobApplication) error {
	if err := services.DB.Create(jobApplication).Error; err != nil {
		return errors.Wrap(err, "failed to create job application")
	}
	return nil
}

// DeleteJobApplication removes a job application from the database by ID
func DeleteJobApplication(id int) error {
	if err := services.DB.Delete(&JobApplication{}, id).Error; err != nil {
		return errors.Wrap(err, "failed to delete job application")
	}
	return nil
}

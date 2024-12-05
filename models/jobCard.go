package models

import (
	"github.com/2deadmen/domestic_backend/services"
	"github.com/pkg/errors"
)

// CreateJobCard inserts a new JobCard into the database.
func CreateJobCard(jobCard *JobCard) error {
	if err := services.DB.Create(jobCard).Error; err != nil {
		return errors.Wrap(err, "failed to create job card")
	}
	return nil
}

// GetAllJobCards retrieves all JobCards from the database.
func GetAllJobCards() ([]JobCard, error) {
	var jobCards []JobCard
	if err := services.DB.Find(&jobCards).Error; err != nil {
		return nil, errors.Wrap(err, "failed to fetch job cards")
	}
	return jobCards, nil
}

// GetJobCardByID retrieves a JobCard by its ID.
func GetJobCardByID(id int) (*JobCard, error) {
	var jobCard JobCard
	if err := services.DB.First(&jobCard, "id = ?", id).Error; err != nil {
		return nil, errors.Wrap(err, "failed to fetch job card")
	}
	return &jobCard, nil
}

// UpdateJobCardByID updates the details of a JobCard by its ID.
func UpdateJobCardByID(id int, updatedJobCard *JobCard) error {
	var jobCard JobCard
	if err := services.DB.First(&jobCard, "id = ?", id).Error; err != nil {
		return errors.Wrap(err, "job card not found")
	}

	if err := services.DB.Model(&jobCard).Updates(updatedJobCard).Error; err != nil {
		return errors.Wrap(err, "failed to update job card")
	}
	return nil
}

// DeleteJobCardByID deletes a JobCard by its ID.
func DeleteJobCardByID(id int) error {
	if err := services.DB.Delete(&JobCard{}, "id = ?", id).Error; err != nil {
		return errors.Wrap(err, "failed to delete job card")
	}
	return nil
}

// GetActiveJobCards retrieves all active JobCards from the database.
func GetActiveJobCards() ([]JobCard, error) {
	var jobCards []JobCard
	if err := services.DB.Where("active = ?", true).Find(&jobCards).Error; err != nil {
		return nil, errors.Wrap(err, "failed to fetch active job cards")
	}
	return jobCards, nil
}

// UpdateJobCardActiveStatusByID updates the active status of a JobCard by its ID.
func UpdateJobCardActiveStatusByID(id int, active bool) error {
	var jobCard JobCard
	if err := services.DB.First(&jobCard, "id = ?", id).Error; err != nil {
		return errors.Wrap(err, "job card not found")
	}

	if err := services.DB.Model(&jobCard).Update("active", active).Error; err != nil {
		return errors.Wrap(err, "failed to update job card active status")
	}
	return nil
}

package models

import (
	"github.com/2deadmen/domestic_backend/services"
	"github.com/pkg/errors"
	"time"
)

type JobCard struct {
	Id                      int       `json:"id" gorm:"primaryKey"`
	Pincode                 int       `json:"pincode"`
	Location                string    `json:"location"`
	Gender                  string    `json:"gender"`
	JobType                 []string  `json:"jobType" gorm:"type:text"`
	Salary                  string    `json:"salary"`
	Duration                string    `json:"duration"`
	ExperienceReq           string    `json:"experienceReq"`
	EmployementAvailability time.Time `json:"employementAvailability"`
	WorkingHours            string    `json:"workingHours"`
	Holidays                string    `json:"holidays"`
	EmployerId              int       `json:"employerId" gorm:"foreignKey"`
	Vacancy                 int       `json:"vacancy"`
	Active                  bool      `json:"active"`
}

type JobApplication struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	EmployerId int    `json:"employerId" gorm:"foreignKey"`
	EmployeeId int    `json:"employeeId" gorm:"foreignKey"`
	JobId      int    `json:"jobId" gorm:"foreignKey"`
	Status     string `json:"status"` // "accepted" or "rejected"
}

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

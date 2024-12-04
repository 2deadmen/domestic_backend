package models

import (
	"log"

	"github.com/2deadmen/domestic_backend/services"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

// CreateJobCard inserts a new JobCard into the database.

func CreateJobCard(jobCard *JobCard) error {
	// Debug: Print jobType for verification
	log.Printf("JobType: %v", jobCard.JobType)

	// Prepare query with pq.Array for jobType
	query := `
		INSERT INTO job_cards 
		(pincode, location, gender, job_type, salary, duration, experience_req, employement_availability, working_hours, holidays, employer_id, vacancy, active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id`

	// Using pq.Array to serialize the jobType slice to text[]
	if err := services.DB.Exec(query,
		jobCard.Pincode,                 // pincode
		jobCard.Location,                // location
		jobCard.Gender,                  // gender
		pq.Array(jobCard.JobType),       // job_type as TEXT[]
		jobCard.Salary,                  // salary
		jobCard.Duration,                // duration
		jobCard.ExperienceReq,           // experience_req
		jobCard.EmployementAvailability, // employement_availability
		jobCard.WorkingHours,            // working_hours
		jobCard.Holidays,                // holidays
		jobCard.EmployerId,              // employer_id
		jobCard.Vacancy,                 // vacancy
		jobCard.Active,                  // active
	).Error; err != nil {
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

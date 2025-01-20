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

func GetApplicationsByEmployer(employerId string) ([]JobApplication, error) {
	var applications []JobApplication
	if err := services.DB.Where("employer_id = ?", employerId).Find(&applications).Error; err != nil {
		return nil, errors.Wrap(err, "failed to fetch applications by employer")
	}
	return applications, nil
}

func GetApplicationsByEmployee(employeeId string) ([]JobApplication, error) {
	var applications []JobApplication
	if err := services.DB.Where("employee_id = ?", employeeId).Find(&applications).Error; err != nil {
		return nil, errors.Wrap(err, "failed to fetch applications by employee")
	}
	return applications, nil
}

// GetJobApplicationByID retrieves a job application by its ID
func GetJobApplicationByID(id int) (JobApplication, error) {
	var jobApplication JobApplication
	if err := services.DB.First(&jobApplication, id).Error; err != nil {
		return jobApplication, errors.Wrap(err, "failed to fetch job application")
	}
	return jobApplication, nil
}

// CreateRating adds a new rating to the database
func CreateRating(rating *Rating) error {
	if err := services.DB.Create(rating).Error; err != nil {
		return errors.Wrap(err, "failed to create rating")
	}
	return nil
}

// GetRatingsByEmployee retrieves all ratings and comments for a specific employee
func GetRatingsByEmployee(employeeID uint) ([]Rating, error) {
	var ratings []Rating
	if err := services.DB.Where("employee_id = ?", employeeID).Find(&ratings).Error; err != nil {
		return nil, errors.Wrap(err, "failed to fetch ratings for employee")
	}
	return ratings, nil
}

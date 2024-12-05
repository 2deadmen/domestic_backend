package models

import (
	"time"
)

type Employer struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Email        string    `json:"email" gorm:"unique"`
	Password     string    `json:"password"`
	Age          int       `json:"age"`
	Gender       string    `json:"gender"`
	Phone        string    `json:"phone" `
	AddressProof string    `json:"addressproof"`
	Type         string    `json:"type"`
	OTP          string    `json:"otp"`
	Verified     bool      `json:"verified" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type Employee struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Pin            string    `json:"pin" `
	Age            int       `json:"age"`
	Gender         string    `json:"gender"`
	Phone          string    `json:"phone" `
	AddressProof   string    `json:"addressproof"`
	OpenToWork     bool      `json:"opentowork"`
	WorkExperience string    `json:"workexperience"`
	TypeOfWork     string    `json:"typeofwork"`
	PhotoURL       string    `json:"photourl"`
	Verified       bool      `json:"verified" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type JobCard struct {
	Id                      int       `json:"id" gorm:"primaryKey"`
	Pincode                 int       `json:"pincode"`
	Location                string    `json:"location"`
	Gender                  string    `json:"gender"`
	JobType                 string    `json:"jobType" `
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
	Status     string `json:"status"  gorm:"default:'accepted'"` // "accepted" or "rejected"
}

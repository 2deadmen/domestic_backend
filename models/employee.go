package models

import (
	"github.com/2deadmen/domestic_backend/services"
	"github.com/pkg/errors"
)

// CreateEmployee inserts a new employee into the database
func CreateEmployee(employee *Employee) error {
	if err := services.DB.Create(employee).Error; err != nil {
		return errors.Wrap(err, "failed to create employee")
	}
	return nil
}

// GetEmployeeByPhone retrieves an employee by their phone number
func GetEmployeeByPhone(phone string) (Employee, error) {
	var employee Employee
	err := services.DB.Where("phone = ?", phone).First(&employee).Error
	return employee, err
}

// GetEmployeeByID retrieves an employee by their ID
func GetEmployeeByID(id string) (Employee, error) {
	var employee Employee
	err := services.DB.First(&employee, "id = ?", id).Error
	return employee, err
}

// GetAllEmployees retrieves all employees from the database
func GetAllEmployees() ([]Employee, error) {
	var employees []Employee
	err := services.DB.Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

// UpdateEmployeeByID updates an employee's details by ID
func UpdateEmployeeByID(id string, updatedEmployee *Employee) error {
	var employee Employee
	if err := services.DB.First(&employee, "id = ?", id).Error; err != nil {
		return err
	}

	return services.DB.Model(&employee).Updates(updatedEmployee).Error
}

// DeleteEmployeeByID deletes an employee by their ID
func DeleteEmployeeByID(id string) error {
	return services.DB.Delete(&Employee{}, "id = ?", id).Error
}

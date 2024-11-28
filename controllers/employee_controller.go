package controllers

import (
	"net/http"

	"github.com/2deadmen/domestic_backend/models"
	"github.com/2deadmen/domestic_backend/utils"

	"github.com/gin-gonic/gin"
)

// CreateEmployee godoc
// @Summary Register a new employee
// @Description Sign up a new employee
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body models.Employee true "Employee Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees [post]
func CreateEmployee(c *gin.Context) {
	var employee models.Employee

	// Bind JSON input to the Employee model
	if err := c.ShouldBindJSON(&employee); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	// Validate required fields
	if employee.Name == "" || employee.Phone == "" || employee.Pin == "" {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Name, phone, and pin are required"})
		return
	}

	// Check if the phone number already exists
	if _, err := models.GetEmployeeByPhone(employee.Phone); err == nil {
		utils.RespondJSON(c, http.StatusConflict, gin.H{"error": "Phone number already in use"})
		return
	}

	// Hash the pin before storing it
	hashedPin, err := utils.HashPassword(employee.Pin)
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to hash pin"})
		return
	}
	employee.Pin = hashedPin

	// Mark as verified by default
	employee.Verified = true

	// Save the employee to the database
	if err := models.CreateEmployee(&employee); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Unable to create employee", "details": err.Error()})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(employee.ID, employee.Phone)
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	utils.RespondJSON(c, http.StatusCreated, gin.H{
		"message": "Employee created successfully",
		"phone":   employee.Phone,
		"token":   token,
	})
}

// SignInEmployee godoc
// @Summary Sign in an employee
// @Description Log in an employee with phone and pin
// @Tags Employees
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Phone and Pin"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/sign-in [post]
func SignInEmployee(c *gin.Context) {
	var credentials struct {
		Phone string `json:"phone"`
		Pin   string `json:"pin"`
	}

	// Bind JSON input to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Find the employee by phone
	employee, err := models.GetEmployeeByPhone(credentials.Phone)
	if err != nil {
		utils.RespondJSON(c, http.StatusUnauthorized, gin.H{"error": "Invalid phone"})
		return
	}

	// Verify the provided pin
	if !utils.CheckPasswordHash(credentials.Pin, employee.Pin) {
		utils.RespondJSON(c, http.StatusUnauthorized, gin.H{"error": "Invalid pin"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(employee.ID, employee.Phone)
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with success message and token
	utils.RespondJSON(c, http.StatusOK, gin.H{
		"message": "Sign-in successful",
		"id":      employee.ID,
		"token":   token,
	})
}

// GetAllEmployees godoc
// @Summary Get all employees
// @Description Retrieve a list of all employees
// @Tags Employees
// @Produce json
// @Success 200 {array} models.Employee
// @Failure 500 {object} map[string]string
// @Router /employees [get]
func GetAllEmployees(c *gin.Context) {
	employees, err := models.GetAllEmployees()
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, employees)
}

// GetEmployee godoc
// @Summary Get an employee by ID
// @Description Retrieve an employee's details
// @Tags Employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} models.Employee
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id} [get]
func GetEmployee(c *gin.Context) {
	id := c.Param("id")

	employee, err := models.GetEmployeeByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to fetch employee"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, employee)
}

// UpdateEmployee godoc
// @Summary Update an employee
// @Description Modify an employee's details
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body models.Employee true "Updated Employee Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id} [put]
func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var updatedEmployee models.Employee

	if err := c.ShouldBindJSON(&updatedEmployee); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := models.UpdateEmployeeByID(id, &updatedEmployee); err != nil {
		if err.Error() == "record not found" {
			utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Employee updated successfully"})
}

// DeleteEmployee godoc
// @Summary Delete an employee
// @Description Remove an employee by ID
// @Tags Employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id} [delete]
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	if err := models.DeleteEmployeeByID(id); err != nil {
		if err.Error() == "record not found" {
			utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

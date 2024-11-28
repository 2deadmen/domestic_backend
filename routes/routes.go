package routes

import (
	"github.com/2deadmen/domestic_backend/controllers"
	_ "github.com/2deadmen/domestic_backend/docs" // Import the generated docs
	"github.com/2deadmen/domestic_backend/models"

	// "github.com/2deadmen/domestic_backend/middlewares"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files" // Swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRoutes godoc
// @Summary Initialize API routes
// @Description Sets up all the API routes, including Swagger documentation and user endpoints
// @Tags Routes
func InitRoutes() *gin.Engine {
	router := gin.Default()

	// Migrate the database models to the database
	models.MigrateModels()

	//middleware

	// Swagger route
	// @Summary Swagger UI
	// @Description Serve Swagger API documentation
	// @Tags Documentation
	// @Router /swagger/*any [get]
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	// router.Use(middlewares.JWTMiddleware())

	// Employer routes
	employerGroup := router.Group("/employers")
	{
		// @Summary Register a new employer
		// @Description Add a new employer with email and password
		// @Tags Employers
		// @Accept json
		// @Produce json
		// @Param employer body models.Employer true "Employer Data"
		// @Success 201 {object} map[string]interface{}
		// @Failure 400 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /employers [post]

		employerGroup.POST("/", controllers.CreateEmployer)

		// VerifyOTP godoc
		// @Summary Verify OTP for an employer
		// @Description Validate the OTP sent to the employer's email and activate the account
		// @Tags Employers
		// @Accept json
		// @Produce json
		// @Param otp body models.VerifyOTPRequest true "OTP Verification Data"
		// @Success 200 {object} map[string]interface{}
		// @Failure 400 {object} map[string]string
		// @Failure 404 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /employers/verify-otp [post]
		employerGroup.POST("/verify-otp", controllers.VerifyOTP)

		// SignIn godoc
		// @Summary Sign in an employer
		// @Description Authenticate an employer using email and password, and return a JWT
		// @Tags Employers
		// @Accept json
		// @Produce json
		// @Param credentials body models.SignInRequest true "Sign In Credentials"
		// @Success 200 {object} map[string]interface{}
		// @Failure 400 {object} map[string]string
		// @Failure 401 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /employers/sign-in [post]
		employerGroup.POST("/sign-in", controllers.SignIn)

		// GetAllEmployers godoc
		// @Summary Get all employers
		// @Description Retrieve a list of all employers
		// @Tags Employers
		// @Produce json
		// @Success 200 {array} models.Employer
		// @Failure 500 {object} map[string]string
		// @Router /employers [get]
		employerGroup.GET("/", controllers.GetAllEmployers)

		// @Summary Get an employer by ID
		// @Description Retrieve an employer's details using their ID
		// @Tags Employers
		// @Produce json
		// @Param id path int true "Employer ID"
		// @Success 200 {object} models.Employer
		// @Failure 404 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /employers/{id} [get]

		employerGroup.GET("/:id", controllers.GetEmployer)

		// @Summary Update an employer by ID
		// @Description Modify an employer's details using their ID
		// @Tags Employers
		// @Accept json
		// @Produce json
		// @Param id path int true "Employer ID"
		// @Param employer body models.Employer true "Updated Employer Data"
		// @Success 200 {object} map[string]string
		// @Failure 400 {object} map[string]string
		// @Failure 404 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /employers/{id} [put]
		employerGroup.PUT("/:id", controllers.UpdateEmployer)

		// @Summary Delete an employer by ID
		// @Description Remove an employer from the database using their ID
		// @Tags Employers
		// @Produce json
		// @Param id path int true "Employer ID"
		// @Success 200 {object} map[string]string
		// @Failure 404 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /employers/{id} [delete]
		employerGroup.DELETE("/:id", controllers.DeleteEmployer)

	}

	//    --------------------------------------------------------------------------------------------------------------------------------

	// Employee routes group
	employeeRoutes := router.Group("/employees")
	{
		// @Summary Register a new employee
		// @Description Sign up a new employee
		// @Tags Employees
		// @Accept json
		// @Produce json
		// @Param employee body models.Employee true "Employee Data"
		// @Success 201 {object} map[string]interface{}
		// @Failure 400 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /employees [post]
		employeeRoutes.POST("", controllers.CreateEmployee)

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
		employeeRoutes.POST("/sign-in", controllers.SignInEmployee)

		// @Summary Get all employees
		// @Description Retrieve a list of all employees
		// @Tags Employees
		// @Produce json
		// @Success 200 {array} models.Employee
		// @Failure 500 {object} map[string]string
		// @Router /employees [get]
		employeeRoutes.GET("", controllers.GetAllEmployees)

		// @Summary Get an employee by ID
		// @Description Retrieve an employee's details
		// @Tags Employees
		// @Produce json
		// @Param id path int true "Employee ID"
		// @Success 200 {object} models.Employee
		// @Failure 404 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /employees/{id} [get]
		employeeRoutes.GET(":id", controllers.GetEmployee)

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
		employeeRoutes.PUT(":id", controllers.UpdateEmployee)

		// @Summary Delete an employee
		// @Description Remove an employee by ID
		// @Tags Employees
		// @Produce json
		// @Param id path int true "Employee ID"
		// @Success 200 {object} map[string]string
		// @Failure 404 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /employees/{id} [delete]
		employeeRoutes.DELETE(":id", controllers.DeleteEmployee)
	}

	return router
}

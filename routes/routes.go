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
	// User routes
	userGroup := router.Group("/users")
	{
		// @Summary Get all users
		// @Description Retrieve a list of all users
		// @Tags Users
		// @Produce json
		// @Router /users [get]
		userGroup.GET("/", controllers.GetUsers)

		// @Summary Create a new user
		// @Description Add a new user to the system
		// @Tags Users
		// @Accept json
		// @Produce json
		// @Router /users [post]
		userGroup.POST("/", controllers.CreateUser)
	}

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

	return router
}

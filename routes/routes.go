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

	return router
}

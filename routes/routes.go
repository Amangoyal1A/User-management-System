package routes

import (
	"myapp/controllers"
	"myapp/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Inject DB into context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/users", controllers.GetAllUsers)
		protected.GET("/users/:id", controllers.GetUserByID)
		protected.PUT("/users/:id", controllers.UpdateUser)
		protected.DELETE("/users/:id", controllers.DeleteUser)
	}

	return r
}

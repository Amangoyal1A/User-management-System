package main

import (
	"user-management/config"
	"user-management/connection"
	"user-management/controllers"
	"user-management/repository"
	"user-management/routes"
	"user-management/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration and initialize database
	cfg := config.LoadConfig()
	db := connection.InitDB(cfg)

	// Dependency injection
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userControllers := controllers.NewUserController(userService)

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, userControllers)

	// Run server
	r.Run(":8080")
}

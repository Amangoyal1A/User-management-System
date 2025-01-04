package main

import (
	"log"
	"myapp/config"
	"myapp/routes"
)

func main() {
	// Initialize the database
	db := config.InitDB()

	// Ensure the underlying SQL DB connection is closed when the app exits
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}
	defer sqlDB.Close()

	// Set up and run routes
	r := routes.SetupRoutes(db)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

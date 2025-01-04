package config

import (
	"fmt"
	"log"
	"myapp/model"
	"myapp/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("myapp.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate models
	db.AutoMigrate(&models.User{}, &models.Post{})

	// Seed Data
	SeedData(db)

	return db
}

func SeedData(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
		db.Create(&models.User{Name: "Admin", Email: "admin@example.com",  Password: utils.HashPassword("password")})
		fmt.Println("Seeded default admin user")
	}
}

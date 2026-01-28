package main

import (
	"log"
	"tender-app-be-go/internal/db"
	"tender-app-be-go/internal/models"
	"tender-app-be-go/internal/response"
	"tender-app-be-go/internal/seed"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := db.Open("./internal/db/data/app.db")
	if err != nil {
		log.Fatal((err))
	}

	// migrate tables
	if err := db.AutoMigrate(
		&models.User{},
		&models.Vendor{},
		&models.Tender{},
		&models.TenderVendor{},
		&models.TenderProduct{},
	); err != nil {
		log.Fatal(err)
	}

	// seed default data
	if err := seed.Ensure(db); err != nil {
		log.Fatal(err)
	}

	api := r.Group("/api")
	api.GET("/version", func(c *gin.Context) {
		response.OK(c, "v1.0-local", 1, "API for Test Connection")
	})

	// Endpoint to get all users
	api.GET("/users", func(c *gin.Context) {
		var users []models.User // Create a slice to hold the data

		// This is the "SELECT * FROM users" query
		if err := db.Find(&users).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Use your response helper to show the data
		response.OK(c, users, len(users), "Get all users success")
	})

	r.Run(":8000")
}

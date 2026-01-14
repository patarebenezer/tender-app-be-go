package main

import (
	"log"
	"tender-app-be-go/internal/db"
	"tender-app-be-go/internal/models"
	"tender-app-be-go/internal/response"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := db.Open("./internal/db/data/app.db")
	if err != nil {
		log.Fatal((err))
	}
	db.AutoMigrate(&models.User{}, &models.Vendor{}, &models.Tender{}, &models.TenderVendor{}, &models.TenderProduct{})
	log.Println("SQLite Connected!")
	_ = db

	api := r.Group("/api")
	api.GET("/version", func(c *gin.Context) {
		response.OK(c, "v1.0-local", 1, "API for Test Connection")
	})

	r.Run(":8000")
}

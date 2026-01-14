package main

import (
	"tender-app-be-go/internal/response"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")                     // this root
	api.GET("/version", func(c *gin.Context) { // this child : /api/version
		response.OK(c, "v1.0-local", 1, "API for Test Connection") // response
	})

	r.Run(":8000")
}

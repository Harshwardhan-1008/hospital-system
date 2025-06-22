package main

import (
	"net/http"

	"hospital-system/config"
	"hospital-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	routes.SetupRoutes(r)

	// Home route to confirm deployment is live
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hospital Management System is running ✅"})
	})

	// Start server on port 8080 (Render uses this)
	r.Run(":8080")
}

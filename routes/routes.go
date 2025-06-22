package routes

import (
	"hospital-system/controllers"

	"hospital-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Login route
	router.POST("/login", controllers.Login)

	// Health check route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Add patient route
	router.POST("/patients", middleware.AuthMiddleware("doctor", "receptionist"), controllers.AddPatient)
	router.GET("/patients", middleware.AuthMiddleware("doctor", "receptionist"), controllers.GetAllPatients)
	router.GET("/patients/:id", middleware.AuthMiddleware("doctor", "receptionist"), controllers.GetPatientByID)
	router.PUT("/patients/:id", middleware.AuthMiddleware("doctor"), controllers.UpdatePatient)
	router.DELETE("/patients/:id", middleware.AuthMiddleware("doctor"), controllers.DeletePatient)

}

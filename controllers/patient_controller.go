package controllers

import (
	"hospital-system/config"
	"hospital-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPatient(c *gin.Context) {
	var patient models.Patient

	// Bind JSON body to struct
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save to DB
	result := config.DB.Create(&patient)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add patient"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Patient added successfully", "data": patient})
}
func GetAllPatients(c *gin.Context) {
	var patients []models.Patient

	// Get all patients from DB
	result := config.DB.Find(&patients)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}
func GetPatientByID(c *gin.Context) {
	id := c.Param("id") // Get ID from URL
	var patient models.Patient

	result := config.DB.First(&patient, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}
func UpdatePatient(c *gin.Context) {
	id := c.Param("id") // Get ID from URL
	var patient models.Patient

	// Step 1: Check if patient exists
	result := config.DB.First(&patient, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Step 2: Bind new data to the existing patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Step 3: Save updates to DB
	config.DB.Save(&patient)
	c.JSON(http.StatusOK, gin.H{"message": "Patient updated successfully", "data": patient})
}
func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	// Step 1: Check if patient exists
	result := config.DB.First(&patient, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Step 2: Delete the patient
	config.DB.Delete(&patient)
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}

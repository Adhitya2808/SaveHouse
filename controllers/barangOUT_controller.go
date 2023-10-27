package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"github.com/labstack/echo/v4"
	"time"
	"net/http"
)

func CreateBarangOUT(c echo.Context) error {
	// Parse JSON request body into a BarangIN models
	var barangkeluar models.BarangOUT
	if barangkeluar.Transaction_OUT.IsZero() {
		barangkeluar.Transaction_OUT = time.Now()
	}
	if err := c.Bind(&barangkeluar); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	// Create the new BarangIN in the database
	if err := config.DB.Create(&barangkeluar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create Barang_IN"})
	}

	// Return a JSON response with the created BarangIN
	return c.JSON(http.StatusCreated, barangkeluar)
}

func GetAllBarangOUT(c echo.Context) error {
	var barangs []models.BarangOUT

	if err := config.DB.Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve room types"})
	}

	return c.JSON(http.StatusOK, barangs)
}

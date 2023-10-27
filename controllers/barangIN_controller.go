package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
)

func CreateBarangIN(c echo.Context) error {
	// Parse JSON request body into a BarangIN models
	var Barangmasuk models.BarangIN
	if Barangmasuk.Transaction_IN.IsZero() {
		Barangmasuk.Transaction_IN = time.Now()
	}
	if err := c.Bind(&Barangmasuk); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	// Create the new BarangIN in the database
	if err := config.DB.Create(&Barangmasuk).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create Barang_IN"})
	}
	// Return a JSON response with the created BarangIN
	return c.JSON(http.StatusCreated, Barangmasuk)
}

func GetAllBarangIN(c echo.Context) error {
	var barangs []models.BarangIN

	if err := config.DB.Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve room types"})
	}

	return c.JSON(http.StatusOK, barangs)
}

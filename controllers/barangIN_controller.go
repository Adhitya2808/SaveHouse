package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateBarangIN(c echo.Context) error {
	// Parse JSON request body into a BarangIN models
	var Barangmasuk models.BarangIN
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

func DeleteBarangIN(c echo.Context) error {
	id := c.Param("trx_id")

	var barang models.BarangIN
	if err := config.DB.First(&barang, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve Barang_IN"})
	}

	if err := config.DB.Delete(&barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete Barang_IN"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Barang_IN successfully deleted"})
}

func UpdateBarangIN(c echo.Context) error {
	id := c.Param("trx_id")

	var barang models.BarangIN
	if err := config.DB.First(&barang, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve room type"})
	}

	if err := c.Bind(&barang).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot bind tipe kamar"})
	}

	if err := config.DB.Save(&barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update room type"})
	}

	return c.JSON(http.StatusOK, barang)

}
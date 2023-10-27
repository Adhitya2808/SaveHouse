package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateBarang(c echo.Context) error {
	// Parse JSON request body into a Barang models
	Barang := models.Barang{}
	if err := c.Bind(&Barang); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Create the new Barang in the database
	if err := config.DB.Create(&Barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create Barang_IN"})
	}

	// Return a JSON response with the created BarangIN
	return c.JSON(http.StatusCreated, Barang)
}

func GetAllBarang(c echo.Context) error {
	var barangs []models.Barang

	if err := config.DB.Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve room types"})
	}

	return c.JSON(http.StatusOK, barangs)
}

func DeleteBarang(c echo.Context) error {
	id := c.Param("id")

	var barang models.BarangIN
	if err := config.DB.First(&barang, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve Barang_IN"})
	}

	if err := config.DB.Delete(&barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete Barang_IN"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Barang_IN successfully deleted"})
}

func UpdateBarang(c echo.Context) error {
	id := c.Param("id")

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
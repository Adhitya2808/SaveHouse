package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateTipeGudang(c echo.Context) error {
	// Parse JSON request body into a Barang models
	var gudang models.TipeGudang

	if err := c.Bind(&gudang); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Create the new Barang in the database
	if err := config.DB.Create(&gudang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create TipeGudang Detail"})
	}
	// Return a JSON response with the created Barang
	return c.JSON(http.StatusCreated, gudang)
}

func UpdateGudang(c echo.Context) error {
	id := c.Param("id_gudang")

	var gudang models.TipeGudang
	if err := config.DB.First(&gudang, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve TipeGudang"})
	}

	if err := c.Bind(&gudang).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot bind TipeGudang data"})
	}

	if err := config.DB.Save(&gudang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update TipeGudang"})
	}

	return c.JSON(http.StatusOK, gudang)

}

func DeleteTipeGudang(c echo.Context) error {
	id := c.Param("id_gudang")

	var gudang models.TipeGudang
	if err := config.DB.First(&gudang, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve TipeGudang"})
	}

	if err := config.DB.Delete(&gudang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete TipeGudang"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "TipeGudang successfully deleted"})
}

func GetAllTipeGudang(c echo.Context) error {
	var gudang []models.TipeGudang

	if err := config.DB.Find(&gudang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve All TipeGudang"})
	}

	return c.JSON(http.StatusOK, gudang)
}

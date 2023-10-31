package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllBarangIN(c echo.Context) error {
	var barangs []models.BarangIN

	if err := config.DB.Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve Barang IN"})
	}

	return c.JSON(http.StatusOK, barangs)
}
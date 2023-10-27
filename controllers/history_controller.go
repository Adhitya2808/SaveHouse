package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllHistory(c echo.Context) error {
	var barangs []models.Barang
	if err := config.DB.Preload("Barangmasuk").Preload("Barangkeluar").Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve History Barang"})
	}
	return c.JSON(http.StatusOK, barangs)
}
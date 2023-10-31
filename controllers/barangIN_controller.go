package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetAllBarangIN(c echo.Context) error {
	var barangs []models.BarangIN

	if err := config.DB.Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve room types"})
	}

	return c.JSON(http.StatusOK, barangs)
}
package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"SaveHouse/service"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetAllHistory(c echo.Context) error {
	var barang []models.Barang
	if err := config.DB.Preload("Barangmasuk").Preload("Barangkeluar").Find(&barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve History Barang"})
	}
	var responselist []models.HistoryResponse
	for _, barang := range barang {
		response := service.AllHistoryResponse(barang)
		responselist = append(responselist, response)
	}
	return c.JSON(http.StatusOK, responselist)
}

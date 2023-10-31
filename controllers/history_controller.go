package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"SaveHouse/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetAllHistory(c echo.Context) error {
	var barang []models.Barang
	if err := config.DB.Preload("Barangmasuk").Preload("Barangkeluar").Find(&barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve History Barang"})
	}
	var responselist []models.HistoryResponse
	for _, barang := range barang {
		response := utils.AllHistoryResponse(barang)
		responselist = append(responselist, response)
	}
	return c.JSON(http.StatusOK, responselist)
}

func Searching(c echo.Context) error {
	var barang []models.Barang
	search := c.QueryParam("search")
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	offset := (pageInt - 1) * limitInt

	if err := config.DB.Preload("Barangmasuk").Preload("Barangkeluar").Where("barang_name LIKE ?", "%"+search+"%").Offset(offset).Limit(limitInt).Find(&barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve Barang"})
	}

	if search == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to retrieve Barang"})
	}

	var responselist []models.HistoryResponse
	for _, Barang := range barang {
		response := utils.AllHistoryResponse(Barang)
		responselist = append(responselist, response)
	}
	return c.JSON(http.StatusOK, responselist)

}

package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"SaveHouse/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllBarangIN(c echo.Context) error {
	var barangs []models.BarangIN

	if err := config.DB.Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve BarangIN"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("BarangIN data successfully retrieved", barangs))
}

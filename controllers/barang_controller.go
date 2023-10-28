package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"SaveHouse/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateBarang(c echo.Context) error {
	// Parse JSON request body into a Barang models
	var Barang models.Barang

	if err := c.Bind(&Barang); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Create the new Barang in the database
	fileheader := "photo"
	Barang.Photo = service.CloudinaryUpload(c, fileheader)
	if err := config.DB.Create(&Barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create Barang Detail"})
	}

	var BarangMasuk models.BarangIN
	BarangMasuk.Transaction_IN = config.DB.NowFunc()
	BarangMasuk.Trx_id = Barang.ID
	if err := config.DB.Create(&BarangMasuk).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create Barang_IN"})
	}
	if err := config.DB.Preload("Barangmasuk").Find(&Barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve Barang"})
	}
	models.BarangResponseConvert(Barang)

	// Return a JSON response with the created Barang
	return c.JSON(http.StatusCreated, Barang)
}

func GetBarangByID(c echo.Context) error {
	id := c.Param("id")
	var barang []models.Barang
	if err := config.DB.First(&barang, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve Barang"})
	}
	var responselist []models.BarangResponse
	for _, Barang := range barang {
		response := service.AllBarangsResponse(Barang)
		responselist = append(responselist, response)
	}
	return c.JSON(http.StatusOK, responselist)
}

func UpdateBarang(c echo.Context) error {
	id := c.Param("id")

	var barang models.Barang
	if err := config.DB.First(&barang, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve BarangID"})
	}

	if err := c.Bind(&barang).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot bind BarangID"})
	}

	if err := config.DB.Save(&barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update BarangID"})
	}

	return c.JSON(http.StatusOK, barang)

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

func GetAllBarang(c echo.Context) error {
	var barangs []models.Barang

	if err := config.DB.Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve room types"})
	}

	return c.JSON(http.StatusOK, barangs)
}

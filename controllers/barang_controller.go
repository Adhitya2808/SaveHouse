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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	var barang models.Barang
	if err := c.Bind(&barang); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot bind BarangID"})
	}

	var existingBarang models.Barang
	result := config.DB.First(&existingBarang, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve BarangID"})
	}
	existingBarang.Barang_Name = barang.Barang_Name
	existingBarang.TipeGudang = barang.TipeGudang
	existingBarang.Category = barang.Category
	existingBarang.Description = barang.Description
	existingBarang.Photo = barang.Photo
	existingBarang.Quantity = barang.Quantity
	config.DB.Model(&existingBarang).Updates(barang)

	response := utils.AllBarangsResponse(existingBarang)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Barang data successfully updated", response))

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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve All Barang"})
	}

	return c.JSON(http.StatusOK, barangs)
}

package controllers

import (
	"SaveHouse/config"
	"SaveHouse/models"
	"SaveHouse/service"
	"SaveHouse/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateBarang(c echo.Context) error {
	// Parse JSON request body into a Barang models
	var Barang models.Barang

	if err := c.Bind(&Barang); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	// Create the new Barang in the database
	fileheader := "photo"
	Barang.Photo = service.CloudinaryUpload(c, fileheader)
	if err := config.DB.Create(&Barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create Barang"))
	}

	var BarangMasuk models.BarangIN
	BarangMasuk.Transaction_IN = config.DB.NowFunc()
	BarangMasuk.Trx_id = Barang.ID
	if err := config.DB.Create(&BarangMasuk).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create Barang"))
	}
	if err := config.DB.Preload("Barangmasuk").Find(&Barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve Barang"))
	}
	models.BarangResponseConvert(Barang)

	// Return a JSON response with the created Barang
	return c.JSON(http.StatusCreated, Barang)
}

func GetBarangByID(c echo.Context) error {
	id := c.Param("id")
	var barang []models.Barang
	if err := config.DB.First(&barang, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve Barang"))
	}
	var responselist []models.BarangResponse
	for _, Barang := range barang {
		response := utils.AllBarangsResponse(Barang)
		responselist = append(responselist, response)
	}
	return c.JSON(http.StatusOK, responselist)
}

func UpdateBarang(c echo.Context) error {
	var updatedBarang models.Barang
	if err := c.Bind(&updatedBarang); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("cannot bind barang"))
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Invalid ID"))
	}
	existingbarang := models.Barang{}
	err = config.DB.First(&existingbarang, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve Barang"))
	}
	fileheader := "photo"
	existingbarang.Photo = service.CloudinaryUpload(c, fileheader)
	updatedBarang.Photo = service.CloudinaryUpload(c, fileheader)

	existingbarang.Barang_Name = updatedBarang.Barang_Name
	existingbarang.TipeGudang = updatedBarang.TipeGudang
	existingbarang.Photo = updatedBarang.Photo
	existingbarang.Quantity = updatedBarang.Quantity
	existingbarang.Category = updatedBarang.Category
	existingbarang.Description = updatedBarang.Description

	err = config.DB.Save(&existingbarang).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update Barang"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Barang data successfully updated", utils.AllBarangsResponse(existingbarang)))

}

func DeleteBarang(c echo.Context) error {
	id := c.Param("id")

	var barang models.BarangIN
	if err := config.DB.First(&barang, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve Barang"))
	}

	if err := config.DB.Delete(&barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete Barang"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Barang data successfully deleted", barang))
}

func GetAllBarang(c echo.Context) error {
	var barangs []models.Barang

	if err := config.DB.Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve Barang"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success Retrieve Data", barangs))
}

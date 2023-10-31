package controllers

import (
	"app/config"
	"app/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateBarangOUT(c echo.Context) error {
	// Parse JSON request body into a Barang models
	IdBarang, _ := strconv.Atoi(c.QueryParam("id_barang"))

	var BarangKeluar models.BarangOUT
	var Barang models.Barang
	BarangKeluar.Trx_id = uint(IdBarang)
	BarangKeluar.Transaction_OUT = config.DB.NowFunc()

	// Create the new Barang in the database
	if err := config.DB.Create(&BarangKeluar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create Barang Detail"})
	}
	Barang.ID = uint(IdBarang)
	if err := config.DB.Preload("Barangmasuk").Preload("Barangkeluar").Find(&Barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve Barang"})
	}
	models.BarangResponseConvert(Barang)

	// Return a JSON response with the created Barang
	return c.JSON(http.StatusCreated, Barang)
}

func GetAllBarangOUT(c echo.Context) error {
	var barangs []models.BarangOUT

	if err := config.DB.Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve all barang out"})
	}

	return c.JSON(http.StatusOK, barangs)
}

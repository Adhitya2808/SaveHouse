package models

import (
	"gorm.io/gorm"
	"time"
)

type Barang struct {
	gorm.Model
	Barang_Name  string    `json:"name_of_goods" form:"nama_barang"`
	Category     string    `json:"category" form:"category"`
	Description  string    `json:"description" form:"description"`
	TipeGudang   string    `json:"tipe_gudang" form:"tipe_gudang"`
	Photo        string    `json:"photo" form:"photo"`
	Quantity     int       `json:"qty" form:"qty"`
	Barangmasuk  BarangIN  `gorm:"foreignKey:Trx_id"`
	Barangkeluar BarangOUT `gorm:"foreignKey:Trx_id"`
}

type BarangOUT struct {
	Trx_id          uint      `json:"id_barang" form:"trx_id"`
	Transaction_OUT time.Time `json:"transaction_out" form:"transaction_out"`
}

type BarangIN struct {
	Trx_id         uint      `json:"id_barang" form:"trx_id"`
	Transaction_IN time.Time `json:"transaction_in" form:"transaction_in"`
}

type BarangResponse struct {
	Trx_id      uint   `json:"trx_id"`
	Barang_Name string `json:"name_of_goods"`
	Category    string `json:"category"`
	TipeGudang  string `json:"tipe_gudang"`
	Description string `json:"description"`
	Quantity    int    `json:"qty"`
	Photo       string `json:"photo"`
	BarangIN    struct {
		Transaction_IN time.Time `json:"transaction_in"`
	} `json:"barang_in"`
	BarangOUT struct {
		Transaction_OUT time.Time `json:"transaction_out"`
	} `json:"barang_out"`
}

func BarangResponseConvert(barang Barang) BarangResponse {
	var barangResponse BarangResponse

	barangResponse.Barang_Name = barang.Barang_Name
	barangResponse.Category = barang.Category
	barangResponse.Description = barang.Description
	barangResponse.TipeGudang = barang.TipeGudang
	barangResponse.Quantity = barang.Quantity
	barangResponse.Photo = barang.Photo
	barangResponse.BarangIN.Transaction_IN = barang.Barangmasuk.Transaction_IN
	barangResponse.BarangOUT.Transaction_OUT = barang.Barangkeluar.Transaction_OUT
	return barangResponse
}

type HistoryResponse struct {
	Trx_id      uint   `json:"trx_id"`
	Barang_Name string `json:"name_of_goods"`
	Category    string `json:"category"`
	Description string `json:"description"`
	TipeGudang  string `json:"tipe_gudang"`
	Quantity    int    `json:"qty"`
	Photo       string `json:"photo"`
	BarangIN    struct {
		Transaction_IN time.Time `json:"transaction_in"`
	} `json:"barang_in"`
	BarangOUT struct {
		Transaction_OUT time.Time `json:"transaction_out"`
	} `json:"barang_out"`
}

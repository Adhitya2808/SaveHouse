package models

import (
	"time"
	"gorm.io/gorm"
)

type Barang struct {
	gorm.Model
	Trx_id       uint      `json:"trx_id" form:"trx_id"`
	Barang_Name  string    `json:"nama_barang" form:"nama_barang"`
	Category     string    `json:"category" form:"category"`
	Description  string    `json:"description" form:"description"`
	Photo        string    `json:"photo" form:"photo"`
	Quantity     int       `json:"qty" form:"qty"`
	Barangmasuk  BarangIN  `gorm:"foreignKey:Trx_id"`
	Barangkeluar BarangOUT `gorm:"foreignKey:Trx_id"`
}

type BarangOUT struct {
	Trx_id          uint   	  `json:"id_barang" form:"trx_id"`
	Transaction_OUT time.Time `gorm:"type:timestamp" json:"transaction_out"`
}

type BarangIN struct {
	Trx_id         uint   	 `json:"id_barang" form:"trx_id"`
	Transaction_IN time.Time `gorm:"type:timestamp" json:"transaction_in"`
}

type BarangResponse struct{
	Trx_id       uint      `json:"trx_id"`
	Barang_Name  string    `json:"nama_barang"`
	Category     string    `json:"category"`
	Description  string    `json:"description"`
	Quantity     int       `json:"qty"`
	Photo		 string	   `json:"photo"`
}

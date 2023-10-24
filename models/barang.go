package models

import "gorm.io/gorm"

type Barang struct{
	gorm.Model
	ID_Barang	uint		`json:"ID_Barang" form:"ID_Barang"`
	Barang_Name string   	`json:"nama_barang" form:"nama_barang"`
	Category    string   	`json:"category" form:"category"`
	Description string   	`json:"description" form:"description"`
	Photo       string 	 	`json:"photo" form:"photo"`
	Quantity    int  		`json:"qty" form:"qty"`
	Barangmasuk   BarangIN  `gorm:"foreignKey:ID_Barang" references:"ID_Barang"`
	Barangkeluar  BarangOUT `gorm:"foreignKey:ID_Barang" references:"ID_Barang"`
}

type BarangOUT struct {
	ID_Barang		uint   `json:"ID_Barang" form:"ID_Barang"`
	Nama_Barang     string `json:"nama_barang" form:"nama_barang"`
	Transaction_OUT string `json:"transaction_out" form:"transaction_out"`
}

type BarangIN struct {
	ID_Barang		uint   `json:"ID_Barang" form:"ID_Barang"`
	Nama_Barang     string `json:"nama_barang" form:"nama_barang"`
	Transaction_IN  string `json:"transaction_in" form:"transaction_in"`
}
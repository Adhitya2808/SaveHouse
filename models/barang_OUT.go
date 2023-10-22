package models

type BarangOUT struct {
	ID					int    `gorm:"primaryKey;type:smallint" json:"id" form:"id"`
	Id_Product			uint   `gorm:"smallint;not null" json:"id_product" form:"id_product"`
	Transaction_OUT		string `json:"transaction_out" form:"transaction_out"`
	Nama_Barang 		string `gorm:"varchar(225);not null" json:"nama_barang" form:"nama_barang"`
}
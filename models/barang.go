package models

import ("time")

type Barang struct {
	ID          int    `gorm:"primaryKey;type:smallint" json:"id" form:"id"`
	Id_Product   int   `gorm:"smallint;not null" json:"id_product" form:"id_product"`
	Nama_Barang string `gorm:"varchar(225);not null" json:"nama_barang" form:"nama_barang"`
	Kategori    string `gorm:"varchar(225);not null" json:"kategori" form:"kategori"`
	Description string `gorm:"type:text;not null" json:"deskripsi" form:"deksripsi"`
	Photo       string `json:"photo" form:"photo"`
	Qty         int32  `gorm:"smallint;not null" json:"qty" form:"qty"`
	Exp_date    time.Time 
	Transaction_INS  []BarangIN `gorm:"foreignKey:Id_Product;references:ID"`
	Transaction_OUTS []BarangOUT `gorm:"foreignKey:Id_Product;references:ID"`
}
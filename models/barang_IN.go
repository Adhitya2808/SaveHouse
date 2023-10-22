package models

type BarangIN struct {
	ID					int    `gorm:"primaryKey;type:smallint" json:"id" form:"id"`
	Id_Product			uint   `gorm:"smallint;not null" json:"id_product" form:"id_product"`
	Nama_Barang 		string `gorm:"varchar(225);not null" json:"nama_barang" form:"nama_barang"`
	Kategori    		string `gorm:"varchar(225);not null" json:"kategori" form:"kategori"`
	Photo				string `json:"photo" form:"photo"`
	Qty					int32  `json:"qty" form:"qty"`
	Transaction_IN		string `json:"transaction_in" form:"transaction_in"`
}
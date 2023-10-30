package models

type GudangRequest struct {
	NamaBarang string `json:"nama_barang" form:"nama_barang"`
	Quantity   string `json:"quantity" form:"quantity"`
}

type AIResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

package utils

import "app/models"

func AllBarangsResponse(barangResponse models.Barang) models.BarangResponse {
	return models.BarangResponse{
		Trx_id:      barangResponse.ID,
		Barang_Name: barangResponse.Barang_Name,
		TipeGudang:  barangResponse.TipeGudang,
		Category:    barangResponse.Category,
		Description: barangResponse.Description,
		Photo:       barangResponse.Photo,
		Quantity:    barangResponse.Quantity,
	}
}

func AllHistoryResponse(historyResponse models.Barang) models.HistoryResponse {
	response := models.HistoryResponse{
		Trx_id:      historyResponse.ID,
		Barang_Name: historyResponse.Barang_Name,
		TipeGudang:  historyResponse.TipeGudang,
		Category:    historyResponse.Category,
		Description: historyResponse.Description,
		Photo:       historyResponse.Photo,
		Quantity:    historyResponse.Quantity,
	}
	response.BarangIN.Transaction_IN = historyResponse.Barangmasuk.Transaction_IN
	response.BarangOUT.Transaction_OUT = historyResponse.Barangkeluar.Transaction_OUT

	return response
}

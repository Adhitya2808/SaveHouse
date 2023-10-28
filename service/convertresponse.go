package service

import "SaveHouse/models"

func AllBarangsResponse(barangResponse models.Barang) models.BarangResponse {
	return models.BarangResponse{
		Barang_Name: barangResponse.Barang_Name,
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
		Category:    historyResponse.Category,
		Description: historyResponse.Description,
		Photo:       historyResponse.Photo,
		Quantity:    historyResponse.Quantity,
	}
	response.BarangIN.Transaction_IN = historyResponse.Barangmasuk.Transaction_IN
	response.BarangOUT.Transaction_OUT = historyResponse.Barangkeluar.Transaction_OUT

	return response
}

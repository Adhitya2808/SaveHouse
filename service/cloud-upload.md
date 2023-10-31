package service

import (
	"SaveHouse/config"
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"os"
)

func UploadImage(imageFile *os.File) (string, error) {
	// Konfigurasi Cloudinary
	config := cloudinary.{
		CloudName: "your_cloud_name",
		APIKey:    "your_api_key",
		APISecret: "your_api_secret",
	}

	// Inisialisasi Cloudinary
	cld, err := cloudinary.NewFromConfiguration(config)
	if err != nil {
		return "photo", err
	}

	// Konfigurasi uploader
	uploaderConfig := uploader.UploadParams{
		Folder: "your_upload_folder", // Nama folder di Cloudinary
	}

	// Upload gambar ke Cloudinary
	uploadResult, err := cld.Upload.Upload(context.Background(), imageFile, uploaderConfig)
	if err != nil {
		return "photo", err
	}

	// Kembalikan URL gambar yang diupload
	return uploadResult.URL, nil
}

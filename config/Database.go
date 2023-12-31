package config

import (
	"SaveHouse/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

type Config struct {
	dbUser        string
	dbPass        string
	dbHost        string
	dbPort        string
	dbName        string
	secret        string
	cloudinaryUrl string
}

func ConnectDB() {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		panic("Failed to Connect Database")
	}

	InitMigrate()

	fmt.Println("Connected to Database")

}

func InitMigrate() {

	DB.AutoMigrate(&models.User{}, &models.Barang{}, &models.BarangIN{}, &models.BarangOUT{})
}

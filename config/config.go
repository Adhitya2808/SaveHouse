package config

import (
	"SaveHouse/models"
	"os"
	"strconv"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
  	DB *gorm.DB
)



type Config struct {
	SERVER_PORT int
	DB_USER string
	DB_PASS string
	DB_HOST	string
	DB_PORT	int
	DB_NAME string
	SIGN_KEY string
	CLOUD_URL string
	REFRESH_KEY string
}



func LoadDBConfig() Config {
	godotenv.Load(".env")

	DB_PORT, err := strconv.Atoi(os.Getenv("DB_PORT"))
	SERVER_PORT, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		panic(err)
	}

	return Config {
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: DB_PORT,
		DB_NAME: os.Getenv("DB_NAME"),
		SERVER_PORT: SERVER_PORT,
		SIGN_KEY: os.Getenv("SIGN_KEY"),
		CLOUD_URL: os.Getenv("CLOUDINARY_URL"),
		REFRESH_KEY: os.Getenv("REFRESH_KEY"),
	}
}



func InitDB() *gorm.DB{

	config := LoadDBConfig()
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitialMigration(db)

	return db
}


func InitialMigration(db *gorm.DB) {
  	db.AutoMigrate(&models.User{},&models.BarangIN{})
	db.AutoMigrate(&models.Barang{},&models.BarangOUT{})
}
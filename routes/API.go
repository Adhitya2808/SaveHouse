package routes

import (
	"SaveHouse/controllers"
	"SaveHouse/middleware"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)


func New() *echo.Echo {

	e := echo.New()

	AdminSecretKey := os.Getenv("ADMIN_SECRET")

	admin := e.Group("")
	admin.Use(echojwt.JWT([]byte(AdminSecretKey)))

	e.Use(middleware.NotFoundHandler)

	//LOGIN USER DAN ADMIN
	e.POST("/login", controllers.UserLogin)
	e.POST("/login/admin", controllers.AdminLogin)

	//BARANG
	e.GET("/history", controllers.GetAllHistory)
	e.GET("/history/", controllers.Searching)
	e.POST("/in/barang", controllers.CreateBarang)
	e.PUT("/in/barang/:id", controllers.UpdateBarang)
	e.DELETE("/in/barang/:id", controllers.DeleteBarang)
	e.POST("/out/barang", controllers.CreateBarangOUT)
	e.GET("/barang", controllers.GetAllBarang)

	//ADMIN
	admin.GET("/akun", controllers.AllUser)
	admin.GET("/akun/", controllers.UserbyID)
	e.POST("/akun", controllers.Store)
	admin.DELETE("/akun/:id", controllers.UserDelete)
	e.PUT("/akun/:id", controllers.UserUpdate)


	e.POST("/rekomendasi", controllers.GetRecommendation)

	return e
}

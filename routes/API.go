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
	UserSecretKey := os.Getenv("USER_SECRET")

	auth := e.Group("")
	auth.Use(echojwt.JWT([]byte(AdminSecretKey)))

	e.Use(middleware.NotFoundHandler)

	//LOGIN USER DAN ADMIN
	e.POST("/login", controllers.UserLogin)
	e.POST("/admin/login", controllers.AdminLogin)

	//BARANG
	user := e.Group("")
	user.Use(echojwt.JWT([]byte(UserSecretKey)))
	e.GET("/history", controllers.GetAllHistory)
	e.GET("/history/", controllers.Searching)
	e.POST("/in/barang", controllers.CreateBarang)
	e.PUT("/in/barang/:id", controllers.UpdateBarang)
	e.DELETE("/in/barang/:id", controllers.DeleteBarang)
	e.POST("/out/barang", controllers.CreateBarangOUT)
	e.GET("/barang", controllers.GetAllBarang)

	//ADMIN
	e.GET("/akun", controllers.AllUser)
	e.GET("/akun/", controllers.UserbyID)
	e.POST("/akun", controllers.UserRegister)
	auth.DELETE("/akun/:id", controllers.UserDelete)
	e.PUT("/akun/:id", controllers.UserUpdate)

	//TIPE GUDANG
	e.GET("/tipegudang", controllers.GetAllTipeGudang)
	e.POST("/tipegudang", controllers.CreateTipeGudang)
	e.PUT("/tipegudang/:id", controllers.UpdateGudang)
	e.DELETE("/tipegudang/:id", controllers.DeleteTipeGudang)

	e.POST("/rekomendasi", controllers.GetRecommendation)

	auth.GET("/jwttester/admin", controllers.JwtTesterAdmin)
	user.GET("/jwttester/user", controllers.JwtTesterUser)

	return e
}

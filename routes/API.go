package routes

import (
	"SaveHouse/controllers"
	"SaveHouse/middleware"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo{
	e := echo.New()

	AdminSecretkey := os.Getenv("ADMIN_SECRET")
	//UserSecretkey	 := os.Getenv("USER_SECRET")
	e.Use(middleware.NotFoundHandler)

	//LOGIN USER DAN ADMIN
	e.POST("admin/login", controllers.LoginAdmin)
	e.POST("user/login", controllers.UserLogin)

	//user := e.Group("")
	//user.Use(jwt.JWT([]byte(Userkey)))
	e.GET("/history", controllers.GetAllHistory)
	e.POST("/barang", controllers.CreateBarang)
	e.GET("/barang", controllers.GetAllBarang)
	e.GET("/barang/:id", controllers.GetBarangByID)
	e.PUT("/barang/:id", controllers.UpdateBarang)
	e.DELETE("/barang/:id", controllers.DeleteBarang)
	
	//barang masuk
	e.POST("/barangmasuk", controllers.CreateBarangIN)
	e.GET("/barangmasuk", controllers.GetAllBarangIN)

	//barang keluar
	e.POST("/barangkeluar", controllers.CreateBarangOUT)
	e.GET("/barangkeluar", controllers.GetAllBarangOUT)

	//ADMIN 
	Admin := e.Group("/admin")
	Admin.Use(echojwt.JWT([]byte(AdminSecretkey)))
	Admin.GET("/akun", controllers.AllUser)
	e.GET("/akun/:id", controllers.UserbyID)
	e.POST("/akun", controllers.UserRegister)
	e.DELETE("/akun/:id", controllers.UserDelete)
	e.PUT("/akun/:id", controllers.UserUpdate)

	
	return e
}
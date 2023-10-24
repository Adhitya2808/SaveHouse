package routes

import (
	//"SaveHouse/constants"
	"SaveHouse/controllers"
	h "SaveHouse/middleware"
	//"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo{
	e := echo.New()
	h.LogMiddleware(e)

	//Login
	e.POST("/login", controllers.LoginAdmin)
	e.POST("/login", controllers.UserLogin)

	//Barang
	//user := e.Group("/user")
	//user.Use(echojwt.JWT([]byte(constants.USER_SECRET_JWT)))
	//e.POST("/barang", controllers.CreateBarangIN)
	//e.GET("/history", controllers.GetHistoryeBarangController)
	e.POST("/barang", controllers.CreateBarang)
	e.GET("/barang", controllers.GetAllBarang)
	e.PUT("/barang/:id", controllers.UpdateBarang)
	e.DELETE("/barang/:id", controllers.DeleteBarang)
	
	//barang masuk
	e.POST("/barangmasuk", controllers.CreateBarangIN)
	e.GET("/barangmasuk", controllers.GetAllBarangIN)
	e.PUT("/barangmasuk/:id", controllers.UpdateBarangIN)
	e.DELETE("/barangmasuk/:id", controllers.DeleteBarangIN)

	//barang keluar
	//e.POST("/barangkeluar", controllers.CreateBarangOUTController)
	//e.GET("/barangkeluar", controllers.GetBarangOUTController)
	//e.GET("/barangkeluar/:id", controllers.GetBarangOUTbyIDController)
	//e.PUT("/barangkeluar/:id", controllers.UpdateBarangOUTController)
	//e.DELETE("/barangkeluar/:id", controllers.DeleteBarangOUTController)

	//User
	//JWT := e.Group("/admin")
	//JWT.Use(echojwt.JWT([]byte(constants.ADMIN_SECRET_JWT)))
	e.GET("/akun", controllers.AllUser)
	e.GET("/akun/:id", controllers.UserbyID)
	e.POST("/akun", controllers.UserRegister)
	e.DELETE("/akun/:id", controllers.UserDelete)
	e.PUT("/akun/:id", controllers.UserUpdate)

	
	return e
}
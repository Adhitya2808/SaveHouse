package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

<<<<<<< Updated upstream
type jwtCustomClaims struct {
=======
type JwtCustomClaims struct {
>>>>>>> Stashed changes
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

<<<<<<< Updated upstream
func CreateTokenUser(userId int, name string) string {
	var payloadParser jwtCustomClaims
	UserSecretKey := os.Getenv("USER_SECRET")

	payloadParser.ID = uint(userId)
	payloadParser.Name = name
	payloadParser.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 60))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payloadParser)
	t, _ := token.SignedString([]byte(UserSecretKey))
	return t
}

func CreateTokenAdmin(userId int, name string) string {
	var payloadParser jwtCustomClaims
=======
func CreateTokenAdmin(userId int, name string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return "gagal mengakses .env"
	}

	var payloadParser JwtCustomClaims
>>>>>>> Stashed changes
	AdminSecretKey := os.Getenv("ADMIN_SECRET")

	payloadParser.ID = uint(userId)
	payloadParser.Name = name
	payloadParser.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 60))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payloadParser)
	t, _ := token.SignedString([]byte(AdminSecretKey))
	return t
}

func HashPassword(password string) string {
	result, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(result)
}

func ComparePassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func NotFoundHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		if err != nil {
			if he, ok := err.(*echo.HTTPError); ok {
				if he.Code == http.StatusNotFound {
					errorMessage := "Invalid Endpoint"
					return c.JSON(http.StatusNotFound, map[string]interface{}{
						"message": errorMessage,
					})
				}
			}

			fmt.Println("Terjadi kesalahan:", err)
		}

		return err
	}
}
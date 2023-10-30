package controllers

import "github.com/labstack/echo/v4"

func JwtTesterAdmin(c echo.Context) error {
	return c.String(200, "JWT Tester Admin")
}
func JwtTesterUser(c echo.Context) error {
	return c.String(200, "JWT Tester User")
}

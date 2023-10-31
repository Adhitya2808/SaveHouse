package controllers

import (
	"app/config"
	"app/middleware"
	"app/models"
	"app/models/web"
	"app/utils"
	"app/utils/req"
	"app/utils/res"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UserbyID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid User ID"))
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrive user"))
	}

	response := res.ConvertGeneral(&user)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully retrieve", response))
}

func Store(c echo.Context) error {
	var user web.UserRequest
	user.Role = models.UserRole

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	userDb := req.PassBody(user)

	// Hash the user's password before storing it
	userDb.Password = middleware.HashPassword(userDb.Password)

	if err := config.DB.Create(&userDb).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store user data"))
	}

	// Return the response without including a JWT token
	response := res.ConvertGeneral(userDb)

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", response))
}

func UserLogin(c echo.Context) error {
	var loginRequest web.UserLoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}
	var user models.User
	if err := config.DB.Where("username = ? AND role = ?", loginRequest.Username, models.UserRole).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credential"))
	}

	if err := middleware.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credential"))
	}
	response := web.UserLoginResponse{
		Username: user.Username,
		Role:     string(user.Role),
		Name:     user.Name,
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("Login User Successful", response))
}

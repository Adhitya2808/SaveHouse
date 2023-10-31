package controllers

import (
	"SaveHouse/config"
	"SaveHouse/middleware"
	"SaveHouse/models"
	"SaveHouse/models/web"
	"SaveHouse/utils"
	"SaveHouse/utils/res"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func AllUser(c echo.Context) error {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndex(users)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully retrieved", response))
}

func UserUpdate(c echo.Context) error {
	var updateuser models.User
	if err := c.Bind(&updateuser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed tp retrieve userid"})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	existingUser := models.User{}
	err = config.DB.First(&existingUser, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed tp retrieve userid"})
	}
	existingUser.Name = updateuser.Name
	existingUser.Username = updateuser.Username
	existingUser.Password = middleware.HashPassword(updateuser.Password)

	err = config.DB.Save(&existingUser).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update user"})
	}

	response := res.ConvertGeneral(&existingUser)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully updated", response))
}

func UserDelete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var existingUser models.User
	result := config.DB.First(&existingUser, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	config.DB.Delete(&existingUser)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully deleted", nil))
}

func AdminLogin(c echo.Context) error {
	var loginRequest web.UserLoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var user models.User
	if err := config.DB.Where("username = ? AND role = ?", loginRequest.Username, models.AdminRole).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Role is not Admin"))
	}

	if err := middleware.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
	}

	token := middleware.CreateTokenAdmin(int(user.ID), user.Name)

	// Buat respons dengan data yang diminta
	response := web.AdminLoginResponse{
		Username: user.Username,
		Role:     string(user.Role),
		Name:     user.Name,
		Token:    token,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("LoginUser successful", response))
}

package controllers

import (
	"SaveHouse/config"
	"SaveHouse/middleware"
	"SaveHouse/models"
	"SaveHouse/models/web"
	"SaveHouse/utils"
	"SaveHouse/utils/req"
	"SaveHouse/utils/res"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func AllUser(c echo.Context) error {
	var users []models.User

	role := middleware.ExtractTokenUserRole(c)
	if role != "Admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Role is not Admin"))
	}
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

	role := middleware.ExtractTokenUserRole(c)
	if role != "Admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Role is not Admin"))
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
	if err := config.DB.Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Role is not Admin"))
	}

	if err := middleware.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
	}

	token, err := middleware.CreateToken(int(user.ID), user.Username, string(user.Role))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create token"))
	}
	// Buat respons dengan data yang diminta
	response := web.AdminLoginResponse{
		Username: user.Username,
		Role:     string(user.Role),
		Name:     user.Name,
		Token:    token,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Login successful", response))
}

func UserbyID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid User ID"))
	}

	role := middleware.ExtractTokenUserRole(c)
	if role != "Admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Role is not Admin"))
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

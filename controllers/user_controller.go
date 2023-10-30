package controllers

import (
	"SaveHouse/config"
	"SaveHouse/middleware"
	"SaveHouse/models"
	"SaveHouse/models/web"
	"SaveHouse/utils"
	"SaveHouse/utils/req"
	"SaveHouse/utils/res"
	"SaveHouse/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

)

func UserbyID(c echo.Context) error {
<<<<<<< Updated upstream
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid User ID"))
    }
=======
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid User ID"))
	}
>>>>>>> Stashed changes

    var  user  models.User

    if err := config.DB.First(&user, id).Error; err != nil{
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrive user"))
    }

    response := res.ConvertGeneral(&user)

    return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully retrieve", response))
}

<<<<<<< Updated upstream
func UserRegister(c echo.Context) error {
    var user web.UserRequest
=======
func Store(c echo.Context) error {
	var user web.UserRequest
	user.Role = models.UserRole
>>>>>>> Stashed changes

    if err := c.Bind(&user); err != nil{
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid Request Body"))
    }

    userDB := req.PassBody(user)
    userDB.Password = middleware.HashPassword(userDB.Password)

    if err := config.DB.Create(&userDB).Error; err != nil{
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("failed to regist user data"))
    }

    response := res.ConvertGeneral(userDB)
    return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", response))
}

func UserLogin(c echo.Context) error {
    var loginRequest web.UserLoginRequest

    if err := c.Bind(&loginRequest); err != nil{
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
    }
    var user models.User
    if err := config.DB.Where("username = ? AND role = user", loginRequest.Username).First(&user).Error; err != nil{
        return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credential"))
    }

<<<<<<< Updated upstream
    if err := middleware.ComparePassword(user.Password, loginRequest.Password); err != nil{
        return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credential"))
    }

    token := middleware.CreateTokenUser(int(user.ID), user.Name)

    response := web.UserLoginResponse{
        Username: user.Username,
        Password: user.Password,
        Token   : token,
    }
    return c.JSON(http.StatusOK, utils.SuccessResponse("Login User Successful", response))
}
=======
	if err := middleware.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credential"))
	}
	response := web.UserLoginResponse{
		Username: user.Username,
		Name:     user.Name,
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("Login User Successful", response))
}
>>>>>>> Stashed changes

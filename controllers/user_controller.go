package controllers

import (
	"SaveHouse/config"
	//h "SaveHouse/helpers"
    //m "SaveHouse/middleware"
	"SaveHouse/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	//"gorm.io/gorm"
)

func GetUsersController(c echo.Context) error {
    var users []models.User
    if err := config.DB.Find(&users).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success get all users",
        "users":   users,
    })
}

func GetUserController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }
    var user []models.User
    if err := config.DB.First(&user, id).Error; err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "User not found")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success get user",
        "user":    user,
    })
}

func CreateUserController(c echo.Context) error {
    user := models.User{}
    c.Bind(&user)
    if err := config.DB.Create(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success create new user",
        "user":    user,
    })
}

func DeleteUserController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }
	var user models.User
    if err := config.DB.Delete(&user, id).Error; err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
    }
    return c.JSON(http.StatusOK, map[string]string{
        "message": "User deleted",
    })
}

func UpdateUserController(c echo.Context) error {
    var update_user models.User
    if err := c.Bind(&update_user); err != nil{
        return c.JSON(http.StatusBadRequest, models.BaseResponse{
            Status: false,
            Message: "User Not Found",
            Data : nil,
        })
    }
    ID, err := strconv.Atoi(c.Param("id"))
    if err != nil{
        return c.JSON(http.StatusBadRequest, models.BaseResponse{
            Status: false,
            Message: "Invalid User ID",
            Data : nil,
        })
    }
    exsitinguser := models.User{}
    err = config.DB.First(&exsitinguser, ID).Error
    if err != nil {
            return c.JSON(http.StatusInternalServerError, models.BaseResponse{
            Status: false,
            Message: "Failed to Fetch user",
            Data : nil,
        })
    }
    exsitinguser.Username = update_user.Username
    exsitinguser.Name = update_user.Name
    exsitinguser.Email = update_user.Email
    exsitinguser.Role = update_user.Role
    exsitinguser.Password = update_user.Password

    err = config.DB.Save(&exsitinguser).Error
    if err != nil{
        return c.JSON(http.StatusInternalServerError, models.BaseResponse{
            Status: false,
            Message: "Failed to update user",
            Data : nil,
        })
    }
    return c.JSON(http.StatusOK, models.BaseResponse{
            Status: true,
            Message: "Update Successfully",
            Data : nil,
    })
}

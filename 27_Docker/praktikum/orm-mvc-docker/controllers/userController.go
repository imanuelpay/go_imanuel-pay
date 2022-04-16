package controllers

import (
	"mvc/config"
	"mvc/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    users,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id " + c.Param("id"),
		"data":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		errMessage := map[string]string{}

		if len(user.Name) < 1 {
			errMessage["name"] = "name is required"
		}

		if len(user.Email) < 1 {
			errMessage["email"] = "email is required"
		}

		if len(user.Password) < 1 {
			errMessage["password"] = "password is required"
		}

		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": errMessage,
		})
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"data":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := config.DB.Model(&models.User{}).Where("id=?", id).Updates(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by id " + c.Param("id"),
		"data":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id " + c.Param("id"),
	})
}

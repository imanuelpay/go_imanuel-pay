package handler

import (
	"go-ca/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase model.UserUseCase
}

func CreateUserHandler(route *echo.Echo, userUseCase model.UserUseCase) {
	userHandler := UserHandler{userUseCase}

	route.GET("/users", userHandler.GetAllUsers)
	route.POST("/users", userHandler.CreateUser)
}

func (u *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := u.userUseCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": users,
	})
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	user := model.User{}

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": err.Error(),
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
			"error": errMessage,
		})
	}

	newUser, err := u.userUseCase.Create(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": newUser,
	})
}

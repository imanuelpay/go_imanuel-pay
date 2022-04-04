package handler

import (
	"go-ca/constant"
	"go-ca/handler/middleware"
	"go-ca/model"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase model.UserUseCase
}

func CreateUserHandler(route *echo.Echo, userUseCase model.UserUseCase) {
	userHandler := UserHandler{userUseCase}

	route.POST("/login", userHandler.Login)
	route.GET("/users", userHandler.GetAllUsers, middleware.JWTMiddleware())
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

func (u *UserHandler) Login(c echo.Context) error {
	request := model.UserLoginRequest{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if request.Email == "" || request.Password == "" {
		errMessage := map[string]string{}

		if len(request.Email) < 1 {
			errMessage["email"] = "email is required"
		}

		if len(request.Password) < 1 {
			errMessage["password"] = "password is required"
		}

		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": errMessage,
		})
	}

	user, err := u.userUseCase.GetByEmailAndPassword(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	claims := jwt.MapClaims{
		"userId": user.ID,
		"name":   user.Name,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString([]byte(constant.SECRET_JWT))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	userResponse := model.UserLoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: jwt,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": userResponse,
	})
}

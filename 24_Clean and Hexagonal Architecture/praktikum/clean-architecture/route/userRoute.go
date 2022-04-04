package route

import (
	"go-ca/handler"
	"go-ca/repository"
	"go-ca/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoute(e *echo.Echo, DB *gorm.DB) {
	userRepository := repository.CreateUserRepository(DB)
	userUseCase := usecase.CreateUserUsecase(userRepository)
	handler.CreateUserHandler(e, userUseCase)
}

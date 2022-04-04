package route

import (
	"go-ca/handler/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(DB *gorm.DB) *echo.Echo {
	route := echo.New()
	middleware.Logger(route)

	UserRoute(route, DB)

	return route
}

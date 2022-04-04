package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(DB *gorm.DB) *echo.Echo {
	route := echo.New()
	UserRoute(route, DB)

	return route
}

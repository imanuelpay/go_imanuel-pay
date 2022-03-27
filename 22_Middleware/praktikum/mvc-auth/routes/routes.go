package routes

import (
	"mvc/constants"
	"mvc/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/login", controllers.LoginUserController)
	e.POST("/users", controllers.CreateUserController)
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)

	auth := e.Group("/")
	auth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	auth.GET("users", controllers.GetUsersController)
	auth.GET("users/:id", controllers.GetUserController)
	auth.PUT("users/:id", controllers.UpdateUserController)
	auth.DELETE("users/:id", controllers.DeleteUserController)

	auth.POST("books", controllers.CreateBookController)
	auth.PUT("books/:id", controllers.UpdateBookController)
	auth.DELETE("books/:id", controllers.DeleteBookController)

	return e
}

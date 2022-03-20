package main

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    users,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	count := 0
	user := User{}

	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			user = users[i]
			count = 1
		}
	}

	if len(users) == 0 || count == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "record not found by id " + c.Param("id"),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id " + c.Param("id"),
		"data":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"data":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	count := 0
	user := User{}
	c.Bind(&user)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}

	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			users[i].Id = id

			if len(user.Name) > 0 {
				users[i].Name = user.Name
			}

			if len(user.Email) > 0 {
				users[i].Email = user.Email
			}

			if len(user.Password) > 0 {
				users[i].Password = user.Password
			}

			count = i + 1
		}
	}

	if len(users) == 0 || count == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "record not found by id " + c.Param("id"),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by id " + c.Param("id"),
		"data":    users[count-1],
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	count := 0
	user := User{}
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			count = 1
			user = users[i]
			users[len(users)-1], users[i] = users[i], users[len(users)-1]
		}
	}

	if len(users) == 0 || count == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "record not found by id " + c.Param("id"),
		})
	}

	users = users[:len(users)-1]
	sort.Slice(users, func(i, j int) bool {
		return users[i].Id < users[j].Id
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id " + c.Param("id"),
		"user":    user,
	})
}

func main() {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	e.Logger.Fatal(e.Start(":8080"))
}

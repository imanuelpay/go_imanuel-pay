package controllers

import (
	"mvc/config"
	"mvc/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"data":    books,
	})
}

func GetBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book := models.Book{}
	if err := config.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id " + c.Param("id"),
		"data":    book,
	})
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if book.Name == "" || book.Author == "" || book.Publisher == "" {
		errMessage := map[string]string{}

		if len(book.Name) < 1 {
			errMessage["name"] = "name is required"
		}

		if len(book.Author) < 1 {
			errMessage["author"] = "author is required"
		}

		if len(book.Publisher) < 1 {
			errMessage["publisher"] = "publisher is required"
		}

		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": errMessage,
		})
	}

	if err := config.DB.Save(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"data":    book,
	})
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book := models.Book{}
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := config.DB.Model(&models.Book{}).Where("id=?", id).Updates(book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book by id " + c.Param("id"),
		"data":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book := models.Book{}
	if err := config.DB.Delete(&book, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book by id " + c.Param("id"),
	})
}

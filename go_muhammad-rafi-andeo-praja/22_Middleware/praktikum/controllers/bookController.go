package controllers

import (
	"net/http"
	"project-structure/lib/database"
	"project-structure/models"

	"github.com/labstack/echo/v4"
)

func GetBooks(c echo.Context) error {
	books, err := database.GetBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to get books",
		})
	}
	return c.JSON(http.StatusOK, books)
}

func GetBook(c echo.Context) error {
	id := c.Param("id")
	book, err := database.GetBook(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to get book",
		})
	}
	return c.JSON(http.StatusOK, book)
}

func CreateBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	if err := database.CreateBook(&book); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create book",
		})
	}
	return c.JSON(http.StatusOK, book)
}

func UpdateBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	if err := database.UpdateBook(&book); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to update book",
		})
	}
	return c.JSON(http.StatusOK, book)
}

func DeleteBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	if err := database.DeleteBook(&book); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete book",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Book deleted",
	})
}

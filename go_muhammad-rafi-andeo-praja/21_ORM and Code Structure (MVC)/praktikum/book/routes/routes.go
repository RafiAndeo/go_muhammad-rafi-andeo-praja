package routes

import (
	"project-structure/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/books", controllers.GetBooks)
	e.GET("/books/:id", controllers.GetBook)
	e.POST("/books", controllers.CreateBook)
	e.PUT("/books", controllers.UpdateBook)
	e.DELETE("/books", controllers.DeleteBook)

	return e
}

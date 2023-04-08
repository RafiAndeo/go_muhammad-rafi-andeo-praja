package routes

import (
	"project-structure/constants"
	"project-structure/controllers"
	"project-structure/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// User routes
	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetUser)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users", controllers.UpdateUser)
	e.DELETE("/users", controllers.DeleteUser)

	// Book routes
	e.GET("/books", controllers.GetBooks)
	e.GET("/books/:id", controllers.GetBook)
	e.POST("/books", controllers.CreateBook)
	e.PUT("/books", controllers.UpdateBook)
	e.DELETE("/books", controllers.DeleteBook)

	eAuth := e.Group("")
	eAuth.Use(echoMid.BasicAuth(middlewares.BasicAuthDB))

	r := eAuth.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users/:id", controllers.GetUserDetailControllers)

	eAuth.DELETE("/users", controllers.DeleteUser)
	eAuth.PUT("/users", controllers.UpdateUser)

	return e
}

package routes

import (
	"project-structure/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// User routes
	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetUser)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users", controllers.UpdateUser)
	e.DELETE("/users", controllers.DeleteUser)

	// Blog routes
	e.GET("/blogs", controllers.GetBlogs)
	e.GET("/blogs/:id", controllers.GetBlog)
	e.POST("/blogs", controllers.CreateBlog)
	e.PUT("/blogs", controllers.UpdateBlog)
	e.DELETE("/blogs", controllers.DeleteBlog)

	return e
}

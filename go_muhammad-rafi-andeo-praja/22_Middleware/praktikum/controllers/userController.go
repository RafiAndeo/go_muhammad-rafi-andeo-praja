package controllers

import (
	"net/http"
	"project-structure/lib/database"
	"project-structure/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := database.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	err := database.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	err := database.UpdateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	err := database.DeleteUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user)
}

func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	users, err := database.LoginUsers(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login Success",
		"users":   users,
	})
}

func GetUserDetailControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := database.GetDetailUsers(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get User Detail Success",
		"user":    user,
	})
}

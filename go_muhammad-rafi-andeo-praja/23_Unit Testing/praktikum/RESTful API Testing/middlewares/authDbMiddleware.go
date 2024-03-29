package middlewares

import (
	"project-structure/config"
	"project-structure/models"

	"github.com/labstack/echo/v4"
)

var db = config.DB

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var user models.User
	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

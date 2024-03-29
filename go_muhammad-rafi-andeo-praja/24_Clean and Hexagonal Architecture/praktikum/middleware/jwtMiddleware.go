package middlewares

import (
	"belajar-go-echo/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenUserId(c echo.Context) (string, error) {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)
		return userId, nil
	}
	return "", nil
}

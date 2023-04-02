package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_gorm",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	var user User

	if err := DB.First(&user, c.Param("id")).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	var user User

	if err := DB.Delete(&user, c.Param("id")).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := User{}
	if err := DB.First(&user, c.Param("id")).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

func main1() {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.GET("/user/:id", GetUserController)
	e.POST("/user", CreateUserController)
	e.DELETE("/user/:id", DeleteUserController)
	e.PUT("/user/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}

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
		DB_Name:     "blogs_crud_gorm",
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
	Blogs    []Blog `gorm:"foreignKey:UserID"`
}

type Blog struct {
	gorm.Model
	UserID uint   `json:"user_id" form:"user_id"`
	Title  string `json:"title" form:"title"`
	Body   string `json:"body" form:"body"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{}, &Blog{})
}

// get all blogs
func GetBlogsController(c echo.Context) error {
	var blogs []Blog

	if err := DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	var blog Blog

	if err := DB.First(&blog, c.Param("id")).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog",
		"blog":    blog,
	})
}

// create new blog
func CreateBlogController(c echo.Context) error {
	blog := Blog{}
	c.Bind(&blog)

	if err := DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	var blog Blog

	if err := DB.Delete(&blog, c.Param("id")).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	blog := Blog{}
	if err := DB.First(&blog, c.Param("id")).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&blog)

	if err := DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog":    blog,
	})
}

func main() {
	e := echo.New()

	e.GET("/blogs", GetBlogsController)
	e.GET("/blogs/:id", GetBlogController)
	e.POST("/blogs", CreateBlogController)
	e.DELETE("/blogs/:id", DeleteBlogController)
	e.PUT("/blogs/:id", UpdateBlogController)

	e.Logger.Fatal(e.Start(":8000"))
}

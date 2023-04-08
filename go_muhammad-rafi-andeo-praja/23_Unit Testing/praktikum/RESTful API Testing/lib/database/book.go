package database

import (
	"project-structure/config"
	"project-structure/models"
)

func GetBook(id string) (models.Book, error) {
	var book models.Book
	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func CreateBook(book *models.Book) error {
	if err := config.DB.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(book *models.Book) error {
	if err := config.DB.Save(&book).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBook(book *models.Book) error {
	if err := config.DB.Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

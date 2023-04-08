package database

import (
	"project-structure/config"
	"project-structure/middlewares"
	"project-structure/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func GetUser(id string) (models.User, error) {
	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user *models.User) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User) error {
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *models.User) error {
	if err := config.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

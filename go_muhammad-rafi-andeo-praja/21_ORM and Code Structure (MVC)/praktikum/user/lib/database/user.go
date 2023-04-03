package database

import (
	"project-structure/config"
	"project-structure/models"
)

func GetUsers() (users []models.User, err error) {
	if err = config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id string) (user models.User, err error) {
	if err = config.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user models.User) (err error) {
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id string) (err error) {
	var user models.User
	if err = config.DB.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(id string, user models.User) (err error) {
	if err = config.DB.First(&user, id).Error; err != nil {
		return err
	}
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

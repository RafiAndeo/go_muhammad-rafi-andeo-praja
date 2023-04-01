package database

import (
	"project/config"
	"project/models"
)

func GetUsers() (users []models.User, err error) {
	if err = config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

package repository

import "belajar-go-echo/model"

type UserRepository interface {
	Fetch() ([]*model.User, error)
	GetById(id int) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Store(user *model.User) (*model.User, error)
	Delete(id int) error
}

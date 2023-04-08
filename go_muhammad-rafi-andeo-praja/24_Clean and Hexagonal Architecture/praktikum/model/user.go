package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique" json:"email"` // unique
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

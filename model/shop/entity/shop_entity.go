package entity

import (
	"time"

	"gorm.io/gorm"
)

type Shop struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password 	string         `json:"password"`
	Description string         `json:"description"`
	Address     string         `json:"address"`
	Status      string         `json:"status"`
	Contact     string         `json:"contact"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdateAt    time.Time      `json:"update_at"`
	DeleteAt    gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
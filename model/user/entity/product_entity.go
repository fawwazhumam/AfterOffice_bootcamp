package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Shop_ID     uint           `json:"shop_id" gorm:"foreignKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       uint           `json:"price"`
	Stoct       uint           `json:"stoct"`
	Category    string         `json:"category"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdateAt    time.Time      `json:"update_at"`
	DeleteAt    gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
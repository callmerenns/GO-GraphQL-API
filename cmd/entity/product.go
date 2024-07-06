package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `gorm:"not null" json:"string"`
	Stock       int            `gorm:"not null" json:"stock"`
	Price       float64        `gorm:"not null" json:"price"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Users       []*User        `gorm:"many2many:enrollments;" json:"users"`
}

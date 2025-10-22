package entity

import (
	"time"

	"gorm.io/gorm"
)

type SideMenu struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	StoreID     uint           `gorm:"not null" json:"store_id"`
	Store       Store          `gorm:"foreignKey:StoreID" json:"store"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Price       *float64       `json:"price"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type CreateSideMenuRequest struct {
	StoreID     uint    `json:"store_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       *float64 `json:"price"`
}

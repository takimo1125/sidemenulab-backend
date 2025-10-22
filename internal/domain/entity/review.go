package entity

import (
	"time"

	"gorm.io/gorm"
)

type SideMenuReview struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	SideMenuID  uint           `gorm:"not null" json:"side_menu_id"`
	SideMenu    SideMenu       `gorm:"foreignKey:SideMenuID" json:"side_menu"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	User        User           `gorm:"foreignKey:UserID" json:"user"`
	Rating      int            `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Title       string         `json:"title"`
	Comment     string         `json:"comment"`
	IsVerified  bool           `gorm:"default:false" json:"is_verified"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type SideMenuReviewImage struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ReviewID    uint      `gorm:"not null" json:"review_id"`
	Review      SideMenuReview `gorm:"foreignKey:ReviewID" json:"review"`
	ImageURL    string    `gorm:"not null" json:"image_url"`
	ImageOrder  int       `gorm:"default:0" json:"image_order"`
	CreatedAt   time.Time `json:"created_at"`
}

type SideMenuReviewLike struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ReviewID  uint      `gorm:"not null" json:"review_id"`
	Review    SideMenuReview `gorm:"foreignKey:ReviewID" json:"review"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateReviewRequest struct {
	SideMenuID uint   `json:"side_menu_id" binding:"required"`
	Rating     int    `json:"rating" binding:"required,min=1,max=5"`
	Title      string `json:"title"`
	Comment    string `json:"comment"`
}

type CreateReviewImageRequest struct {
	ReviewID   uint   `json:"review_id" binding:"required"`
	ImageURL   string `json:"image_url" binding:"required"`
	ImageOrder int    `json:"image_order"`
}

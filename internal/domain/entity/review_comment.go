package entity

import (
	"time"

	"gorm.io/gorm"
)

// ReviewComment レビューコメントエンティティ
type ReviewComment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	ReviewID  uint           `gorm:"not null" json:"review_id"`
	Review    SideMenuReview `gorm:"foreignKey:ReviewID" json:"review"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Comment   string         `gorm:"not null" json:"comment"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// CreateReviewCommentRequest レビューコメント作成リクエスト
type CreateReviewCommentRequest struct {
	ReviewID uint   `json:"review_id" binding:"required"`
	Comment  string `json:"comment" binding:"required"`
}

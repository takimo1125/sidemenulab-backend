package repository

import "sidemenulab-backend/internal/domain/entity"

// ReviewCommentRepository レビューコメントリポジトリインターフェース
type ReviewCommentRepository interface {
	CreateReviewComment(comment *entity.ReviewComment) error
	GetReviewCommentByID(id uint) (*entity.ReviewComment, error)
	GetReviewCommentsByReviewID(reviewID uint) ([]*entity.ReviewComment, error)
	GetReviewCommentsByUserID(userID uint) ([]*entity.ReviewComment, error)
	GetAllReviewComments() ([]*entity.ReviewComment, error)
	UpdateReviewComment(comment *entity.ReviewComment) error
	DeleteReviewComment(id uint) error
}

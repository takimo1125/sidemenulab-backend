package interfaces

import "sidemenulab-backend/internal/domain/entity"

type ReviewCommentUseCase interface {
	CreateReviewComment(req *entity.CreateReviewCommentRequest, userID uint) (*entity.ReviewComment, error)
	GetReviewCommentByID(id uint) (*entity.ReviewComment, error)
	GetReviewCommentsByReviewID(reviewID uint) ([]*entity.ReviewComment, error)
	GetReviewCommentsByUserID(userID uint) ([]*entity.ReviewComment, error)
	GetAllReviewComments() ([]*entity.ReviewComment, error)
	UpdateReviewComment(comment *entity.ReviewComment) error
	DeleteReviewComment(id uint) error
}

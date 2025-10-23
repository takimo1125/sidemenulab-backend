package database

import (
	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/domain/repository"

	"gorm.io/gorm"
)

type ReviewCommentRepository struct {
	db *gorm.DB
}

func NewReviewCommentRepository(db *gorm.DB) repository.ReviewCommentRepository {
	return &ReviewCommentRepository{db: db}
}

func (r *ReviewCommentRepository) CreateReviewComment(comment *entity.ReviewComment) error {
	return r.db.Create(comment).Error
}

func (r *ReviewCommentRepository) GetReviewCommentByID(id uint) (*entity.ReviewComment, error) {
	var comment entity.ReviewComment
	if err := r.db.Preload("Review").Preload("User").First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *ReviewCommentRepository) GetReviewCommentsByReviewID(reviewID uint) ([]*entity.ReviewComment, error) {
	var comments []*entity.ReviewComment
	if err := r.db.Preload("Review").Preload("User").Where("review_id = ?", reviewID).Order("created_at DESC").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *ReviewCommentRepository) GetReviewCommentsByUserID(userID uint) ([]*entity.ReviewComment, error) {
	var comments []*entity.ReviewComment
	if err := r.db.Preload("Review").Preload("User").Where("user_id = ?", userID).Order("created_at DESC").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *ReviewCommentRepository) GetAllReviewComments() ([]*entity.ReviewComment, error) {
	var comments []*entity.ReviewComment
	if err := r.db.Preload("Review").Preload("User").Order("created_at DESC").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *ReviewCommentRepository) UpdateReviewComment(comment *entity.ReviewComment) error {
	return r.db.Save(comment).Error
}

func (r *ReviewCommentRepository) DeleteReviewComment(id uint) error {
	return r.db.Delete(&entity.ReviewComment{}, id).Error
}

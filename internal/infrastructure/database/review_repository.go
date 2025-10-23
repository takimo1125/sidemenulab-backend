package database

import (
	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/domain/repository"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) repository.ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) CreateReview(review *entity.SideMenuReview) error {
	return r.db.Create(review).Error
}

func (r *ReviewRepository) GetReviewByID(id uint) (*entity.SideMenuReview, error) {
	var review entity.SideMenuReview
	if err := r.db.Preload("User").Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("image_order")
	}).First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) GetReviewsByStoreName(storeName string) ([]*entity.SideMenuReview, error) {
	var reviews []*entity.SideMenuReview
	if err := r.db.Preload("User").Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("image_order")
	}).Where("store_name = ?", storeName).Order("created_at DESC").Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewRepository) GetReviewsByUserID(userID uint) ([]*entity.SideMenuReview, error) {
	var reviews []*entity.SideMenuReview
	if err := r.db.Preload("User").Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("image_order")
	}).Where("user_id = ?", userID).Order("created_at DESC").Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewRepository) GetAllReviews() ([]*entity.SideMenuReview, error) {
	var reviews []*entity.SideMenuReview
	if err := r.db.Preload("User").Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("image_order")
	}).Order("created_at DESC").Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewRepository) UpdateReview(review *entity.SideMenuReview) error {
	return r.db.Save(review).Error
}

func (r *ReviewRepository) DeleteReview(id uint) error {
	return r.db.Delete(&entity.SideMenuReview{}, id).Error
}

func (r *ReviewRepository) CreateReviewImage(image *entity.SideMenuReviewImage) error {
	return r.db.Create(image).Error
}

func (r *ReviewRepository) GetReviewImagesByReviewID(reviewID uint) ([]*entity.SideMenuReviewImage, error) {
	var images []*entity.SideMenuReviewImage
	if err := r.db.Where("review_id = ?", reviewID).Order("image_order").Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func (r *ReviewRepository) CreateReviewLike(like *entity.SideMenuReviewLike) error {
	return r.db.Create(like).Error
}

func (r *ReviewRepository) DeleteReviewLike(reviewID uint, userID uint) error {
	return r.db.Where("review_id = ? AND user_id = ?", reviewID, userID).Delete(&entity.SideMenuReviewLike{}).Error
}

func (r *ReviewRepository) GetReviewLikesByReviewID(reviewID uint) ([]*entity.SideMenuReviewLike, error) {
	var likes []*entity.SideMenuReviewLike
	if err := r.db.Preload("User").Where("review_id = ?", reviewID).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}

package repository

import "sidemenulab-backend/internal/domain/entity"

type ReviewRepository interface {
	CreateReview(review *entity.SideMenuReview) error
	GetReviewByID(id uint) (*entity.SideMenuReview, error)
	GetReviewsByStoreName(storeName string) ([]*entity.SideMenuReview, error)
	GetReviewsByUserID(userID uint) ([]*entity.SideMenuReview, error)
	GetAllReviews() ([]*entity.SideMenuReview, error)
	GetLikedReviewsByUserID(userID uint) ([]*entity.SideMenuReview, error)
	UpdateReview(review *entity.SideMenuReview) error
	DeleteReview(id uint) error
	CreateReviewImage(image *entity.SideMenuReviewImage) error
	GetReviewImagesByReviewID(reviewID uint) ([]*entity.SideMenuReviewImage, error)
	DeleteReviewImage(imageID uint) error
	CreateReviewLike(like *entity.SideMenuReviewLike) error
	DeleteReviewLike(reviewID uint, userID uint) error
	GetReviewLikesByReviewID(reviewID uint) ([]*entity.SideMenuReviewLike, error)
}

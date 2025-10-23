package interfaces

import "sidemenulab-backend/internal/domain/entity"

type ReviewUseCase interface {
	CreateReview(req *entity.CreateReviewRequest) (*entity.SideMenuReview, error)
	CreateReviewWithUserID(req *entity.CreateReviewRequest, userID uint) (*entity.SideMenuReview, error)
	GetReviewByID(id uint) (*entity.SideMenuReview, error)
	GetReviewsByStoreName(storeName string) ([]*entity.SideMenuReview, error)
	GetReviewsByUserID(userID uint) ([]*entity.SideMenuReview, error)
	GetAllReviews() ([]*entity.SideMenuReview, error)
	CreateReviewImage(req *entity.CreateReviewImageRequest) (*entity.SideMenuReviewImage, error)
	GetReviewImagesByReviewID(reviewID uint) ([]*entity.SideMenuReviewImage, error)
	CreateReviewLike(reviewID uint, userID uint) (*entity.SideMenuReviewLike, error)
	DeleteReviewLike(reviewID uint, userID uint) error
	GetReviewLikesByReviewID(reviewID uint) ([]*entity.SideMenuReviewLike, error)
}

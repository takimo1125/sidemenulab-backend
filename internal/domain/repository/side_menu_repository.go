package repository

import "sidemenulab-backend/internal/domain/entity"

type StoreRepository interface {
	CreateStore(store *entity.Store) error
	GetStoreByID(id uint) (*entity.Store, error)
	GetAllStores() ([]*entity.Store, error)
	UpdateStore(store *entity.Store) error
	DeleteStore(id uint) error
}

type SideMenuRepository interface {
	CreateSideMenu(sideMenu *entity.SideMenu) error
	GetSideMenuByID(id uint) (*entity.SideMenu, error)
	GetSideMenusByStoreID(storeID uint) ([]*entity.SideMenu, error)
	GetAllSideMenus() ([]*entity.SideMenu, error)
	UpdateSideMenu(sideMenu *entity.SideMenu) error
	DeleteSideMenu(id uint) error
}

type ReviewRepository interface {
	CreateReview(review *entity.SideMenuReview) error
	GetReviewByID(id uint) (*entity.SideMenuReview, error)
	GetReviewsBySideMenuID(sideMenuID uint) ([]*entity.SideMenuReview, error)
	GetReviewsByUserID(userID uint) ([]*entity.SideMenuReview, error)
	GetAllReviews() ([]*entity.SideMenuReview, error)
	UpdateReview(review *entity.SideMenuReview) error
	DeleteReview(id uint) error
	CreateReviewImage(image *entity.SideMenuReviewImage) error
	GetReviewImagesByReviewID(reviewID uint) ([]*entity.SideMenuReviewImage, error)
	CreateReviewLike(like *entity.SideMenuReviewLike) error
	DeleteReviewLike(reviewID uint, userID uint) error
	GetReviewLikesByReviewID(reviewID uint) ([]*entity.SideMenuReviewLike, error)
}

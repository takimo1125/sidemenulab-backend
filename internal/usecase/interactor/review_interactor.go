package interactor

import (
	"fmt"

	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/domain/repository"
	"sidemenulab-backend/internal/usecase/interfaces"
)

type ReviewInteractor struct {
	reviewRepo repository.ReviewRepository
}

func NewReviewInteractor(reviewRepo repository.ReviewRepository) interfaces.ReviewUseCase {
	return &ReviewInteractor{
		reviewRepo: reviewRepo,
	}
}

func (i *ReviewInteractor) CreateReview(req *entity.CreateReviewRequest) (*entity.SideMenuReview, error) {
	// TODO: ユーザー認証からuserIDを取得する必要があります
	// 現在は仮でuserID=1を使用
	userID := uint(1)

	review := &entity.SideMenuReview{
		StoreName:    req.StoreName,
		SideMenuName: req.SideMenuName,
		UserID:       userID,
		Rating:       req.Rating,
		Title:        req.Title,
		Comment:      req.Comment,
		IsVerified:   false, // デフォルトで未確認
	}

	if err := i.reviewRepo.CreateReview(review); err != nil {
		return nil, fmt.Errorf("レビューの作成に失敗しました: %w", err)
	}

	// 作成されたレビューを関連データと一緒に取得
	createdReview, err := i.reviewRepo.GetReviewByID(review.ID)
	if err != nil {
		return nil, fmt.Errorf("作成されたレビューの取得に失敗しました: %w", err)
	}

	return createdReview, nil
}

func (i *ReviewInteractor) CreateReviewWithUserID(req *entity.CreateReviewRequest, userID uint) (*entity.SideMenuReview, error) {
	review := &entity.SideMenuReview{
		StoreName:    req.StoreName,
		SideMenuName: req.SideMenuName,
		UserID:       userID,
		Rating:       req.Rating,
		Title:        req.Title,
		Comment:      req.Comment,
		IsVerified:   false, // デフォルトで未確認
	}

	if err := i.reviewRepo.CreateReview(review); err != nil {
		return nil, fmt.Errorf("レビューの作成に失敗しました: %w", err)
	}

	// 作成されたレビューを関連データと一緒に取得
	createdReview, err := i.reviewRepo.GetReviewByID(review.ID)
	if err != nil {
		return nil, fmt.Errorf("作成されたレビューの取得に失敗しました: %w", err)
	}

	return createdReview, nil
}

func (i *ReviewInteractor) GetReviewByID(id uint) (*entity.SideMenuReview, error) {
	review, err := i.reviewRepo.GetReviewByID(id)
	if err != nil {
		return nil, fmt.Errorf("レビューの取得に失敗しました: %w", err)
	}
	return review, nil
}

func (i *ReviewInteractor) GetReviewsByStoreName(storeName string) ([]*entity.SideMenuReview, error) {
	reviews, err := i.reviewRepo.GetReviewsByStoreName(storeName)
	if err != nil {
		return nil, fmt.Errorf("店舗のレビュー一覧の取得に失敗しました: %w", err)
	}
	return reviews, nil
}

func (i *ReviewInteractor) GetReviewsByUserID(userID uint) ([]*entity.SideMenuReview, error) {
	reviews, err := i.reviewRepo.GetReviewsByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("ユーザーのレビュー一覧の取得に失敗しました: %w", err)
	}
	return reviews, nil
}

func (i *ReviewInteractor) GetAllReviews() ([]*entity.SideMenuReview, error) {
	reviews, err := i.reviewRepo.GetAllReviews()
	if err != nil {
		return nil, fmt.Errorf("レビュー一覧の取得に失敗しました: %w", err)
	}
	return reviews, nil
}

func (i *ReviewInteractor) GetLikedReviewsByUserID(userID uint) ([]*entity.SideMenuReview, error) {
	reviews, err := i.reviewRepo.GetLikedReviewsByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("ユーザーがいいねしたレビュー一覧の取得に失敗しました: %w", err)
	}
	return reviews, nil
}

func (i *ReviewInteractor) CreateReviewImage(req *entity.CreateReviewImageRequest) (*entity.SideMenuReviewImage, error) {
	image := &entity.SideMenuReviewImage{
		ReviewID:   req.ReviewID,
		ImageURL:   req.ImageURL,
		ImageOrder: req.ImageOrder,
	}

	if err := i.reviewRepo.CreateReviewImage(image); err != nil {
		return nil, fmt.Errorf("レビュー画像の作成に失敗しました: %w", err)
	}

	return image, nil
}

func (i *ReviewInteractor) GetReviewImagesByReviewID(reviewID uint) ([]*entity.SideMenuReviewImage, error) {
	images, err := i.reviewRepo.GetReviewImagesByReviewID(reviewID)
	if err != nil {
		return nil, fmt.Errorf("レビュー画像一覧の取得に失敗しました: %w", err)
	}
	return images, nil
}

func (i *ReviewInteractor) CreateReviewLike(reviewID uint, userID uint) (*entity.SideMenuReviewLike, error) {
	like := &entity.SideMenuReviewLike{
		ReviewID: reviewID,
		UserID:   userID,
	}

	if err := i.reviewRepo.CreateReviewLike(like); err != nil {
		return nil, fmt.Errorf("レビューのイイネに失敗しました: %w", err)
	}

	return like, nil
}

func (i *ReviewInteractor) DeleteReviewLike(reviewID uint, userID uint) error {
	if err := i.reviewRepo.DeleteReviewLike(reviewID, userID); err != nil {
		return fmt.Errorf("レビューのイイネ取り消しに失敗しました: %w", err)
	}
	return nil
}

func (i *ReviewInteractor) GetReviewLikesByReviewID(reviewID uint) ([]*entity.SideMenuReviewLike, error) {
	likes, err := i.reviewRepo.GetReviewLikesByReviewID(reviewID)
	if err != nil {
		return nil, fmt.Errorf("レビューのイイネ一覧の取得に失敗しました: %w", err)
	}
	return likes, nil
}

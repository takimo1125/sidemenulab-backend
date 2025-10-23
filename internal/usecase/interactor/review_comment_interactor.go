package interactor

import (
	"fmt"

	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/domain/repository"
	"sidemenulab-backend/internal/usecase/interfaces"
)

type ReviewCommentInteractor struct {
	reviewCommentRepo repository.ReviewCommentRepository
}

func NewReviewCommentInteractor(reviewCommentRepo repository.ReviewCommentRepository) interfaces.ReviewCommentUseCase {
	return &ReviewCommentInteractor{
		reviewCommentRepo: reviewCommentRepo,
	}
}

func (i *ReviewCommentInteractor) CreateReviewComment(req *entity.CreateReviewCommentRequest, userID uint) (*entity.ReviewComment, error) {
	comment := &entity.ReviewComment{
		ReviewID: req.ReviewID,
		UserID:   userID,
		Comment:  req.Comment,
	}

	if err := i.reviewCommentRepo.CreateReviewComment(comment); err != nil {
		return nil, fmt.Errorf("レビューコメントの作成に失敗しました: %w", err)
	}

	// 作成されたコメントを関連データと一緒に取得
	createdComment, err := i.reviewCommentRepo.GetReviewCommentByID(comment.ID)
	if err != nil {
		return nil, fmt.Errorf("作成されたレビューコメントの取得に失敗しました: %w", err)
	}

	return createdComment, nil
}

func (i *ReviewCommentInteractor) GetReviewCommentByID(id uint) (*entity.ReviewComment, error) {
	comment, err := i.reviewCommentRepo.GetReviewCommentByID(id)
	if err != nil {
		return nil, fmt.Errorf("レビューコメントの取得に失敗しました: %w", err)
	}
	return comment, nil
}

func (i *ReviewCommentInteractor) GetReviewCommentsByReviewID(reviewID uint) ([]*entity.ReviewComment, error) {
	comments, err := i.reviewCommentRepo.GetReviewCommentsByReviewID(reviewID)
	if err != nil {
		return nil, fmt.Errorf("レビューコメント一覧の取得に失敗しました: %w", err)
	}
	return comments, nil
}

func (i *ReviewCommentInteractor) GetReviewCommentsByUserID(userID uint) ([]*entity.ReviewComment, error) {
	comments, err := i.reviewCommentRepo.GetReviewCommentsByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("ユーザーのレビューコメント一覧の取得に失敗しました: %w", err)
	}
	return comments, nil
}

func (i *ReviewCommentInteractor) GetAllReviewComments() ([]*entity.ReviewComment, error) {
	comments, err := i.reviewCommentRepo.GetAllReviewComments()
	if err != nil {
		return nil, fmt.Errorf("全レビューコメント一覧の取得に失敗しました: %w", err)
	}
	return comments, nil
}

func (i *ReviewCommentInteractor) UpdateReviewComment(comment *entity.ReviewComment) error {
	if err := i.reviewCommentRepo.UpdateReviewComment(comment); err != nil {
		return fmt.Errorf("レビューコメントの更新に失敗しました: %w", err)
	}
	return nil
}

func (i *ReviewCommentInteractor) DeleteReviewComment(id uint) error {
	if err := i.reviewCommentRepo.DeleteReviewComment(id); err != nil {
		return fmt.Errorf("レビューコメントの削除に失敗しました: %w", err)
	}
	return nil
}

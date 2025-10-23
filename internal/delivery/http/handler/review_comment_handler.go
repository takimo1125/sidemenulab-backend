package handler

import (
	"net/http"
	"strconv"

	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

type ReviewCommentHandler struct {
	reviewCommentUseCase interfaces.ReviewCommentUseCase
}

func NewReviewCommentHandler(reviewCommentUseCase interfaces.ReviewCommentUseCase) *ReviewCommentHandler {
	return &ReviewCommentHandler{
		reviewCommentUseCase: reviewCommentUseCase,
	}
}

// CreateReviewComment レビューコメント作成
func (h *ReviewCommentHandler) CreateReviewComment(c *gin.Context) {
	var req entity.CreateReviewCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 認証されたユーザーIDを取得
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証情報が取得できません"})
		return
	}

	comment, err := h.reviewCommentUseCase.CreateReviewComment(&req, userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "レビューコメントが作成されました", "data": comment})
}

// GetReviewCommentByID レビューコメント詳細取得
func (h *ReviewCommentHandler) GetReviewCommentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	comment, err := h.reviewCommentUseCase.GetReviewCommentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "レビューコメントが見つかりません"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// GetReviewCommentsByReviewID レビュー別コメント一覧取得
func (h *ReviewCommentHandler) GetReviewCommentsByReviewID(c *gin.Context) {
	reviewIDStr := c.Param("reviewId")
	reviewID, err := strconv.ParseUint(reviewIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なレビューIDです"})
		return
	}

	comments, err := h.reviewCommentUseCase.GetReviewCommentsByReviewID(uint(reviewID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// GetReviewCommentsByUserID ユーザー別コメント一覧取得
func (h *ReviewCommentHandler) GetReviewCommentsByUserID(c *gin.Context) {
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なユーザーIDです"})
		return
	}

	comments, err := h.reviewCommentUseCase.GetReviewCommentsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// GetAllReviewComments 全コメント一覧取得
func (h *ReviewCommentHandler) GetAllReviewComments(c *gin.Context) {
	comments, err := h.reviewCommentUseCase.GetAllReviewComments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// UpdateReviewComment レビューコメント更新
func (h *ReviewCommentHandler) UpdateReviewComment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	var req entity.ReviewComment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 認証されたユーザーIDを取得
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証情報が取得できません"})
		return
	}

	// 既存のコメントを取得してユーザーIDを確認
	existingComment, err := h.reviewCommentUseCase.GetReviewCommentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "レビューコメントが見つかりません"})
		return
	}

	// コメントの所有者かどうかを確認
	if existingComment.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "このコメントを編集する権限がありません"})
		return
	}

	req.ID = uint(id)
	req.UserID = userID.(uint)
	req.ReviewID = existingComment.ReviewID

	if err := h.reviewCommentUseCase.UpdateReviewComment(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "レビューコメントが更新されました"})
}

// DeleteReviewComment レビューコメント削除
func (h *ReviewCommentHandler) DeleteReviewComment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	// 認証されたユーザーIDを取得
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証情報が取得できません"})
		return
	}

	// 既存のコメントを取得してユーザーIDを確認
	existingComment, err := h.reviewCommentUseCase.GetReviewCommentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "レビューコメントが見つかりません"})
		return
	}

	// コメントの所有者かどうかを確認
	if existingComment.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "このコメントを削除する権限がありません"})
		return
	}

	if err := h.reviewCommentUseCase.DeleteReviewComment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "レビューコメントが削除されました"})
}

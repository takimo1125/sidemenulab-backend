package handler

import (
	"net/http"
	"strconv"

	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	reviewUseCase interfaces.ReviewUseCase
}

func NewReviewHandler(reviewUseCase interfaces.ReviewUseCase) *ReviewHandler {
	return &ReviewHandler{
		reviewUseCase: reviewUseCase,
	}
}

// CreateReview レビュー作成
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var req entity.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.reviewUseCase.CreateReview(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "レビューが作成されました", "data": review})
}

// GetReviewByID レビュー詳細取得
func (h *ReviewHandler) GetReviewByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	review, err := h.reviewUseCase.GetReviewByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// GetReviewsBySideMenuID サイドメニュー別レビュー一覧取得
func (h *ReviewHandler) GetReviewsBySideMenuID(c *gin.Context) {
	sideMenuIDStr := c.Param("sideMenuId")
	sideMenuID, err := strconv.ParseUint(sideMenuIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なサイドメニューIDです"})
		return
	}

	reviews, err := h.reviewUseCase.GetReviewsBySideMenuID(uint(sideMenuID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// GetAllReviews レビュー一覧取得
func (h *ReviewHandler) GetAllReviews(c *gin.Context) {
	reviews, err := h.reviewUseCase.GetAllReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// CreateReviewImage レビュー画像アップロード
func (h *ReviewHandler) CreateReviewImage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	var req entity.CreateReviewImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ReviewID = uint(id)

	image, err := h.reviewUseCase.CreateReviewImage(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "レビュー画像がアップロードされました", "data": image})
}

// GetReviewImagesByReviewID レビュー画像一覧取得
func (h *ReviewHandler) GetReviewImagesByReviewID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	images, err := h.reviewUseCase.GetReviewImagesByReviewID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": images})
}

// CreateReviewLike レビューにイイネ
func (h *ReviewHandler) CreateReviewLike(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	// TODO: ユーザー認証からuserIDを取得する必要があります
	// 現在は仮でuserID=1を使用
	userID := uint(1)

	like, err := h.reviewUseCase.CreateReviewLike(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "レビューにイイネしました", "data": like})
}

// DeleteReviewLike レビューのイイネ取り消し
func (h *ReviewHandler) DeleteReviewLike(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	// TODO: ユーザー認証からuserIDを取得する必要があります
	// 現在は仮でuserID=1を使用
	userID := uint(1)

	if err := h.reviewUseCase.DeleteReviewLike(uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "レビューのイイネを取り消しました"})
}

// GetReviewLikesByReviewID レビューのイイネ一覧取得
func (h *ReviewHandler) GetReviewLikesByReviewID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	likes, err := h.reviewUseCase.GetReviewLikesByReviewID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": likes})
}

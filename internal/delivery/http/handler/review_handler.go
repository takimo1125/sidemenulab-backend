package handler

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/infrastructure/cloudinary"
	"sidemenulab-backend/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	reviewUseCase     interfaces.ReviewUseCase
	cloudinaryService *cloudinary.CloudinaryService
}

func NewReviewHandler(reviewUseCase interfaces.ReviewUseCase, cloudinaryService *cloudinary.CloudinaryService) *ReviewHandler {
	return &ReviewHandler{
		reviewUseCase:     reviewUseCase,
		cloudinaryService: cloudinaryService,
	}
}

// CreateReview レビュー作成
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var req entity.CreateReviewRequest
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

	fmt.Printf("レビュー作成 - UserID: %d, Request: %+v\n", userID.(uint), req)

	review, err := h.reviewUseCase.CreateReviewWithUserID(&req, userID.(uint))
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

// GetReviewsByStoreName 店舗別レビュー一覧取得
func (h *ReviewHandler) GetReviewsByStoreName(c *gin.Context) {
	storeName := c.Param("storeName")
	if storeName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "店舗名が指定されていません"})
		return
	}

	reviews, err := h.reviewUseCase.GetReviewsByStoreName(storeName)
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

// UploadReviewImages 複数画像アップロード
func (h *ReviewHandler) UploadReviewImages(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	// CloudinaryServiceが利用できない場合のエラーハンドリング
	if h.cloudinaryService == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "画像アップロードサービスが利用できません"})
		return
	}

	// マルチパートフォームを解析
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "マルチパートフォームの解析に失敗しました"})
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "画像ファイルが選択されていません"})
		return
	}

	var uploadedImages []*entity.SideMenuReviewImage
	ctx := context.Background()

	for i, file := range files {
		// ファイル拡張子をチェック
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ファイル %s はサポートされていない形式です", file.Filename)})
			return
		}

		// ファイルサイズをチェック (5MB制限)
		if file.Size > 5*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ファイル %s が大きすぎます（5MB以下にしてください）", file.Filename)})
			return
		}

		// ファイルを開く
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ファイル %s のオープンに失敗しました", file.Filename)})
			return
		}
		defer src.Close()

		// Cloudinaryにアップロード
		timestamp := time.Now().UnixNano()
		publicID := cloudinary.GeneratePublicID(uint(id), timestamp, i)
		folder := cloudinary.GenerateFolderPath()

		uploadResult, err := h.cloudinaryService.UploadImage(ctx, src, folder, publicID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("画像のアップロードに失敗しました: %s", err.Error())})
			return
		}

		// データベースに画像情報を保存
		imageReq := &entity.CreateReviewImageRequest{
			ReviewID:   uint(id),
			ImageURL:   uploadResult.SecureURL, // CloudinaryのセキュアURLを使用
			ImageOrder: i,
		}

		image, err := h.reviewUseCase.CreateReviewImage(imageReq)
		if err != nil {
			// データベース保存に失敗した場合、Cloudinaryからも削除
			h.cloudinaryService.DeleteImage(ctx, uploadResult.PublicID)
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("画像情報の保存に失敗しました: %s", err.Error())})
			return
		}

		uploadedImages = append(uploadedImages, image)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("%d個の画像がアップロードされました", len(uploadedImages)),
		"data":    uploadedImages,
	})
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

	// 認証されたユーザーIDを取得
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証情報が取得できません"})
		return
	}

	like, err := h.reviewUseCase.CreateReviewLike(uint(id), userID.(uint))
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

	// 認証されたユーザーIDを取得
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証情報が取得できません"})
		return
	}

	if err := h.reviewUseCase.DeleteReviewLike(uint(id), userID.(uint)); err != nil {
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

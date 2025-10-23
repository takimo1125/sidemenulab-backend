package http

import (
	"sidemenulab-backend/internal/delivery/http/handler"
	"sidemenulab-backend/internal/delivery/http/middleware"
	"sidemenulab-backend/internal/infrastructure/cloudinary"
	"sidemenulab-backend/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authUseCase interfaces.AuthUseCase, reviewUseCase interfaces.ReviewUseCase, reviewCommentUseCase interfaces.ReviewCommentUseCase, jwtSecret string, cloudinaryService *cloudinary.CloudinaryService) {
	// ハンドラーを初期化
	authHandler := handler.NewAuthHandler(authUseCase)
	reviewHandler := handler.NewReviewHandler(reviewUseCase, cloudinaryService)
	reviewCommentHandler := handler.NewReviewCommentHandler(reviewCommentUseCase)

	// 認証ミドルウェアを初期化
	authMiddleware := middleware.AuthMiddleware(jwtSecret)

	// API v1 グループ
	v1 := r.Group("/api/v1")
	{
		// 認証関連のルート
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authHandler.SignUp)
			auth.POST("/signin", authHandler.SignIn)
			auth.GET("/debug-token", authHandler.DebugToken)
		}

		// レビュー関連のルート
		reviews := v1.Group("/reviews")
		{
			// 認証が必要なルート（より具体的なルートを先に定義）
			reviews.POST("", authMiddleware, reviewHandler.CreateReview)
			reviews.POST("/:id/images", authMiddleware, reviewHandler.CreateReviewImage)
			reviews.PUT("/:id", authMiddleware, reviewHandler.UpdateReview)
			reviews.DELETE("/:id", authMiddleware, reviewHandler.DeleteReview)
			reviews.DELETE("/images/:imageId", authMiddleware, reviewHandler.DeleteReviewImage)
			reviews.POST("/:id/upload-images", authMiddleware, reviewHandler.UploadReviewImages)
			reviews.POST("/:id/like", authMiddleware, reviewHandler.CreateReviewLike)
			reviews.DELETE("/:id/like", authMiddleware, reviewHandler.DeleteReviewLike)
			reviews.GET("/liked", authMiddleware, reviewHandler.GetLikedReviewsByUserID)
			reviews.GET("/store/:storeName", reviewHandler.GetReviewsByStoreName)
			reviews.GET("/:id/images", reviewHandler.GetReviewImagesByReviewID)
			reviews.GET("/:id/likes", reviewHandler.GetReviewLikesByReviewID)

			// 認証が不要なルート（最後に定義）
			reviews.GET("", reviewHandler.GetAllReviews)
			reviews.GET("/:id", reviewHandler.GetReviewByID)
		}

		// レビューコメント関連のルート
		reviewComments := v1.Group("/review-comments")
		{
			// 認証が必要なルート
			reviewComments.POST("", authMiddleware, reviewCommentHandler.CreateReviewComment)
			reviewComments.PUT("/:id", authMiddleware, reviewCommentHandler.UpdateReviewComment)
			reviewComments.DELETE("/:id", authMiddleware, reviewCommentHandler.DeleteReviewComment)

			// 認証が不要なルート（リスト取得のみ）
			reviewComments.GET("", reviewCommentHandler.GetAllReviewComments)
			reviewComments.GET("/:id", reviewCommentHandler.GetReviewCommentByID)
			reviewComments.GET("/review/:reviewId", reviewCommentHandler.GetReviewCommentsByReviewID)
			reviewComments.GET("/user/:userId", reviewCommentHandler.GetReviewCommentsByUserID)
		}
	}
}

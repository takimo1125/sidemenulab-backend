package http

import (
	"sidemenulab-backend/internal/delivery/http/handler"
	"sidemenulab-backend/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authUseCase interfaces.AuthUseCase, sideMenuUseCase interfaces.SideMenuUseCase, reviewUseCase interfaces.ReviewUseCase) {
	// ハンドラーを初期化
	authHandler := handler.NewAuthHandler(authUseCase)
	sideMenuHandler := handler.NewSideMenuHandler(sideMenuUseCase)
	reviewHandler := handler.NewReviewHandler(reviewUseCase)

	// API v1 グループ
	v1 := r.Group("/api/v1")
	{
		// 認証関連のルート
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authHandler.SignUp)
			auth.POST("/signin", authHandler.SignIn)
		}

		// 店舗関連のルート
		stores := v1.Group("/stores")
		{
			stores.POST("", sideMenuHandler.CreateStore)
			stores.GET("", sideMenuHandler.GetAllStores)
			stores.GET("/:id", sideMenuHandler.GetStoreByID)
		}

		// サイドメニュー関連のルート
		sideMenus := v1.Group("/side-menus")
		{
			sideMenus.POST("", sideMenuHandler.CreateSideMenu)
			sideMenus.GET("", sideMenuHandler.GetAllSideMenus)
			sideMenus.GET("/:id", sideMenuHandler.GetSideMenuByID)
			sideMenus.GET("/store/:storeId", sideMenuHandler.GetSideMenusByStoreID)
		}

		// レビュー関連のルート
		reviews := v1.Group("/reviews")
		{
			reviews.POST("", reviewHandler.CreateReview)
			reviews.GET("", reviewHandler.GetAllReviews)
			reviews.GET("/:id", reviewHandler.GetReviewByID)
			reviews.GET("/side-menu/:sideMenuId", reviewHandler.GetReviewsBySideMenuID)
			reviews.POST("/:id/images", reviewHandler.CreateReviewImage)
			reviews.GET("/:id/images", reviewHandler.GetReviewImagesByReviewID)
			reviews.POST("/:id/like", reviewHandler.CreateReviewLike)
			reviews.DELETE("/:id/like", reviewHandler.DeleteReviewLike)
			reviews.GET("/:id/likes", reviewHandler.GetReviewLikesByReviewID)
		}
	}
}

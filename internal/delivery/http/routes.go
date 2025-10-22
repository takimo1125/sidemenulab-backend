package http

import (
	"sidemenulab-backend/internal/delivery/http/handler"
	"sidemenulab-backend/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authUseCase interfaces.AuthUseCase) {
	// 認証ハンドラーを初期化
	authHandler := handler.NewAuthHandler(authUseCase)

	// API v1 グループ
	v1 := r.Group("/api/v1")
	{
		// 認証関連のルート
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authHandler.SignUp)
			auth.POST("/signin", authHandler.SignIn)
		}
	}
}

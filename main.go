package main

import (
	"log"
	"net/http"
	"os"

	deliveryhttp "sidemenulab-backend/internal/delivery/http"
	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/infrastructure/database"
	"sidemenulab-backend/internal/usecase/interactor"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// PostgreSQLデータベース接続
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=postgres user=postgres password=password dbname=sidemenulab port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベース接続に失敗しました:", err)
	}

	// データベース接続をテスト
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("データベース接続の取得に失敗しました:", err)
	}
	defer sqlDB.Close()

	// データベースマイグレーション
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Store{},
		&entity.SideMenu{},
		&entity.SideMenuReview{},
		&entity.SideMenuReviewImage{},
		&entity.SideMenuReviewLike{},
		&entity.ReviewComment{},
	); err != nil {
		log.Fatal("データベースマイグレーションに失敗しました:", err)
	}

	// 初期データの挿入
	if err := database.SeedData(db); err != nil {
		log.Fatal("初期データの挿入に失敗しました:", err)
	}

	// 依存性注入
	userRepo := database.NewUserRepository(db)
	storeRepo := database.NewStoreRepository(db)
	sideMenuRepo := database.NewSideMenuRepository(db)
	reviewRepo := database.NewReviewRepository(db)
	reviewCommentRepo := database.NewReviewCommentRepository(db)
	
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key" // 本番環境では必ず環境変数で設定してください
	}
	authUseCase := interactor.NewAuthInteractor(userRepo, jwtSecret)
	sideMenuUseCase := interactor.NewSideMenuInteractor(storeRepo, sideMenuRepo)
	reviewUseCase := interactor.NewReviewInteractor(reviewRepo)
	reviewCommentUseCase := interactor.NewReviewCommentInteractor(reviewCommentRepo)

	// Ginエンジンの初期化
	engine := gin.Default()
	
	// CORS設定
	engine.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})
	
	// ヘルスチェックエンドポイント
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Sidemenulab Backend API",
			"status":  "running",
		})
	})

	// データベース接続確認エンドポイント
	engine.GET("/health", func(c *gin.Context) {
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "unhealthy",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"database": "connected",
		})
	})

	// ルート設定
	deliveryhttp.SetupRoutes(engine, authUseCase, sideMenuUseCase, reviewUseCase, reviewCommentUseCase, jwtSecret)

	// サーバー起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("サーバーを起動中... ポート: %s", port)
	engine.Run(":" + port)
}
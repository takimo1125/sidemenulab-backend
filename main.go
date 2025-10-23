package main

import (
	"log"
	"net/http"
	"os"

	deliveryhttp "sidemenulab-backend/internal/delivery/http"
	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/infrastructure/cloudinary"
	"sidemenulab-backend/internal/infrastructure/database"
	"sidemenulab-backend/internal/usecase/interactor"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// .envファイルを読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// PostgreSQLデータベース接続
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=password dbname=sidemenulab port=5432 sslmode=disable TimeZone=Asia/Tokyo"
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
	reviewRepo := database.NewReviewRepository(db)
	reviewCommentRepo := database.NewReviewCommentRepository(db)
	
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key" // 本番環境では必ず環境変数で設定してください
	}
	authUseCase := interactor.NewAuthInteractor(userRepo, jwtSecret)
	reviewUseCase := interactor.NewReviewInteractor(reviewRepo)
	reviewCommentUseCase := interactor.NewReviewCommentInteractor(reviewCommentRepo)

	// Cloudinaryサービスの初期化
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	log.Printf("Cloudinary設定 - Cloud Name: %s, API Key: %s, API Secret: %s", 
		cloudName, 
		func() string { if apiKey != "" { return apiKey[:8] + "..." } else { return "未設定" } }(), 
		func() string { if apiSecret != "" { return apiSecret[:8] + "..." } else { return "未設定" } }())

	var cloudinaryService *cloudinary.CloudinaryService
	if cloudName != "" && apiKey != "" && apiSecret != "" {
		var err error
		cloudinaryService, err = cloudinary.NewCloudinaryService(cloudName, apiKey, apiSecret)
		if err != nil {
			log.Printf("Cloudinaryサービスの初期化に失敗しました: %v", err)
			log.Println("Cloudinaryが利用できないため、ローカルファイルアップロードを使用します")
		} else {
			log.Println("Cloudinaryサービスが正常に初期化されました")
		}
	} else {
		log.Println("Cloudinaryの環境変数が設定されていないため、ローカルファイルアップロードを使用します")
	}

	// Ginエンジンの初期化
	engine := gin.Default()
	
	// 静的ファイルの配信設定
	engine.Static("/uploads", "./uploads")
	
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
	deliveryhttp.SetupRoutes(engine, authUseCase, reviewUseCase, reviewCommentUseCase, jwtSecret, cloudinaryService)

	// サーバー起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("サーバーを起動中... ポート: %s", port)
	engine.Run(":" + port)
}
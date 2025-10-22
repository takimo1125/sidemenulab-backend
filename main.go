package main

import (
	"log"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// PostgreSQLデータベース接続
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=postgres user=postgres password= dbname=sidemenulab port=5432 sslmode=disable TimeZone=Asia/Tokyo"
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

	// Ginエンジンの初期化
	engine := gin.Default()
	
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

	// サーバー起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("サーバーを起動中... ポート: %s", port)
	engine.Run(":" + port)
}
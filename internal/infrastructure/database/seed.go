package database

import (
	"log"

	"sidemenulab-backend/internal/domain/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedData 初期データを挿入する
func SeedData(db *gorm.DB) error {
	log.Println("初期データの挿入を開始します...")

	// ユーザーデータの挿入
	if err := seedUsers(db); err != nil {
		return err
	}

	// レビューデータの挿入
	if err := seedReviews(db); err != nil {
		return err
	}

	// レビューコメントデータの挿入
	if err := seedReviewComments(db); err != nil {
		return err
	}

	// レビュー画像データの挿入
	if err := seedReviewImages(db); err != nil {
		return err
	}

	log.Println("初期データの挿入が完了しました")
	return nil
}

// seedUsers ユーザーデータを挿入
func seedUsers(db *gorm.DB) error {
	// 既存のユーザーをチェック
	var count int64
	db.Model(&entity.User{}).Count(&count)
	if count > 0 {
		log.Println("ユーザーデータは既に存在します")
		return nil
	}

	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	users := []entity.User{
		{
			Email:    "admin@sidemenulab.com",
			Password: string(hashedPassword),
			Name:     "管理者",
		},
		{
			Email:    "user1@example.com",
			Password: string(hashedPassword),
			Name:     "田中太郎",
		},
		{
			Email:    "user2@example.com",
			Password: string(hashedPassword),
			Name:     "佐藤花子",
		},
		{
			Email:    "user3@example.com",
			Password: string(hashedPassword),
			Name:     "鈴木一郎",
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	log.Println("ユーザーデータを挿入しました")
	return nil
}

// seedReviews レビューデータを挿入
func seedReviews(db *gorm.DB) error {
	// 既存のレビューをチェック
	var count int64
	db.Model(&entity.SideMenuReview{}).Count(&count)
	if count > 0 {
		log.Println("レビューデータは既に存在します")
		return nil
	}

	reviews := []entity.SideMenuReview{
		{
			UserID:       2,
			StoreName:    "サンプル店舗",
			SideMenuName: "サンプルサイドメニュー",
			Rating:       5,
			Title:        "とても美味しかった！",
			Comment:      "新鮮な野菜で、ドレッシングも絶品でした。また食べたいです。",
			IsVerified:   true,
		},
		{
			UserID:       3,
			StoreName:    "サンプル店舗",
			SideMenuName: "サンプルサイドメニュー",
			Rating:       4,
			Title:        "良いサラダ",
			Comment:      "野菜が新鮮で美味しかったです。",
			IsVerified:   true,
		},
		{
			UserID:       2,
			StoreName:    "サンプル店舗",
			SideMenuName: "サンプルサイドメニュー",
			Rating:       5,
			Title:        "最高のポテトフライ",
			Comment:      "カリッとしていて、塩加減も絶妙でした。",
			IsVerified:   true,
		},
		{
			UserID:       4,
			StoreName:    "サンプル店舗",
			SideMenuName: "サンプルサイドメニュー",
			Rating:       3,
			Title:        "普通のポテトフライ",
			Comment:      "特に特徴はありませんが、美味しかったです。",
			IsVerified:   false,
		},
		{
			UserID:       3,
			StoreName:    "サンプル店舗",
			SideMenuName: "サンプルサイドメニュー",
			Rating:       4,
			Title:        "ジューシーなチキン",
			Comment:      "チキンが柔らかくて美味しかったです。",
			IsVerified:   true,
		},
		{
			UserID:       4,
			StoreName:    "サンプル店舗",
			SideMenuName: "サンプルサイドメニュー",
			Rating:       5,
			Title:        "エビがプリプリ",
			Comment:      "エビが新鮮で、衣もサクサクでした。",
			IsVerified:   true,
		},
	}

	for _, review := range reviews {
		if err := db.Create(&review).Error; err != nil {
			return err
		}
	}

	log.Println("レビューデータを挿入しました")
	return nil
}

// seedReviewComments レビューコメントデータを挿入
func seedReviewComments(db *gorm.DB) error {
	// 既存のコメントをチェック
	var count int64
	db.Model(&entity.ReviewComment{}).Count(&count)
	if count > 0 {
		log.Println("レビューコメントデータは既に存在します")
		return nil
	}

	comments := []entity.ReviewComment{
		{
			ReviewID: 1,
			UserID:   3,
			Comment:  "私も同じサラダを食べました！本当に美味しかったです。",
		},
		{
			ReviewID: 1,
			UserID:   4,
			Comment:  "ドレッシングの種類は何種類ありますか？",
		},
		{
			ReviewID: 3,
			UserID:   3,
			Comment:  "ポテトフライの塩加減が絶妙ですね！",
		},
		{
			ReviewID: 3,
			UserID:   4,
			Comment:  "私も食べてみたいです！",
		},
		{
			ReviewID: 5,
			UserID:   2,
			Comment:  "チキンナゲットも美味しそうですね。",
		},
		{
			ReviewID: 6,
			UserID:   2,
			Comment:  "エビフライ、私も食べてみたいです！",
		},
		{
			ReviewID: 6,
			UserID:   3,
			Comment:  "エビが新鮮で良かったですね。",
		},
	}

	for _, comment := range comments {
		if err := db.Create(&comment).Error; err != nil {
			return err
		}
	}

	log.Println("レビューコメントデータを挿入しました")
	return nil
}

// seedReviewImages レビュー画像データを挿入
func seedReviewImages(db *gorm.DB) error {
	// 既存のレビュー画像をチェック
	var count int64
	db.Model(&entity.SideMenuReviewImage{}).Count(&count)
	if count > 0 {
		log.Println("レビュー画像データは既に存在します")
		return nil
	}

	// テスト用の画像URL（CloudinaryのサンプルURL）
	images := []entity.SideMenuReviewImage{
		{
			ReviewID:   1,
			ImageURL:   "https://res.cloudinary.com/ducbz8w43/image/upload/v1761228134/sidemenulab/reviews/2025/review_9_1761228133545502001_0.jpg",
			ImageOrder: 0,
		},
		{
			ReviewID:   1,
			ImageURL:   "https://res.cloudinary.com/ducbz8w43/image/upload/v1761228134/sidemenulab/reviews/2025/review_9_1761228133545502001_0.jpg",
			ImageOrder: 1,
		},
		{
			ReviewID:   3,
			ImageURL:   "https://res.cloudinary.com/ducbz8w43/image/upload/v1761228134/sidemenulab/reviews/2025/review_9_1761228133545502001_0.jpg",
			ImageOrder: 0,
		},
		{
			ReviewID:   3,
			ImageURL:   "https://res.cloudinary.com/ducbz8w43/image/upload/v1761228134/sidemenulab/reviews/2025/review_9_1761228133545502001_0.jpg",
			ImageOrder: 1,
		},
		{
			ReviewID:   3,
			ImageURL:   "https://res.cloudinary.com/ducbz8w43/image/upload/v1761228134/sidemenulab/reviews/2025/review_9_1761228133545502001_0.jpg",
			ImageOrder: 2,
		},
		{
			ReviewID:   5,
			ImageURL:   "https://res.cloudinary.com/ducbz8w43/image/upload/v1761228134/sidemenulab/reviews/2025/review_9_1761228133545502001_0.jpg",
			ImageOrder: 0,
		},
		{
			ReviewID:   6,
			ImageURL:   "https://res.cloudinary.com/ducbz8w43/image/upload/v1761228134/sidemenulab/reviews/2025/review_9_1761228133545502001_0.jpg",
			ImageOrder: 0,
		},
		{
			ReviewID:   6,
			ImageURL:   "https://res.cloudinary.com/ducbz8w43/image/upload/v1761228134/sidemenulab/reviews/2025/review_9_1761228133545502001_0.jpg",
			ImageOrder: 1,
		},
	}

	for _, image := range images {
		if err := db.Create(&image).Error; err != nil {
			return err
		}
	}

	log.Println("レビュー画像データを挿入しました")
	return nil
}

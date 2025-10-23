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

	// 店舗データの挿入
	if err := seedStores(db); err != nil {
		return err
	}

	// サイドメニューデータの挿入
	if err := seedSideMenus(db); err != nil {
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

// seedStores 店舗データを挿入
func seedStores(db *gorm.DB) error {
	// 既存の店舗をチェック
	var count int64
	db.Model(&entity.Store{}).Count(&count)
	if count > 0 {
		log.Println("店舗データは既に存在します")
		return nil
	}

	stores := []entity.Store{
		{
			Name:    "サイドメニュー研究所 本店",
			Address: "東京都渋谷区道玄坂1-2-3",
			Phone:   "03-1234-5678",
		},
		{
			Name:    "サイドメニュー研究所 新宿店",
			Address: "東京都新宿区新宿3-1-1",
			Phone:   "03-2345-6789",
		},
		{
			Name:    "サイドメニュー研究所 池袋店",
			Address: "東京都豊島区池袋1-1-1",
			Phone:   "03-3456-7890",
		},
	}

	for _, store := range stores {
		if err := db.Create(&store).Error; err != nil {
			return err
		}
	}

	log.Println("店舗データを挿入しました")
	return nil
}

// seedSideMenus サイドメニューデータを挿入
func seedSideMenus(db *gorm.DB) error {
	// 既存のサイドメニューをチェック
	var count int64
	db.Model(&entity.SideMenu{}).Count(&count)
	if count > 0 {
		log.Println("サイドメニューデータは既に存在します")
		return nil
	}

	sideMenus := []entity.SideMenu{
		{
			StoreID:     1,
			Name:        "特製サラダ",
			Description: "新鮮な野菜とドレッシングのサラダ",
			Price:       func() *float64 { v := 580.0; return &v }(),
		},
		{
			StoreID:     1,
			Name:        "ポテトフライ",
			Description: "カリッと揚げたポテトフライ",
			Price:       func() *float64 { v := 450.0; return &v }(),
		},
		{
			StoreID:     1,
			Name:        "オニオンリング",
			Description: "サクサクのオニオンリング",
			Price:       func() *float64 { v := 380.0; return &v }(),
		},
		{
			StoreID:     2,
			Name:        "チキンナゲット",
			Description: "ジューシーなチキンナゲット",
			Price:       func() *float64 { v := 520.0; return &v }(),
		},
		{
			StoreID:     2,
			Name:        "フライドチキン",
			Description: "香ばしく揚げたフライドチキン",
			Price:       func() *float64 { v := 680.0; return &v }(),
		},
		{
			StoreID:     3,
			Name:        "エビフライ",
			Description: "プリプリのエビフライ",
			Price:       func() *float64 { v := 750.0; return &v }(),
		},
		{
			StoreID:     3,
			Name:        "コロッケ",
			Description: "ほくほくのコロッケ",
			Price:       func() *float64 { v := 320.0; return &v }(),
		},
	}

	for _, sideMenu := range sideMenus {
		if err := db.Create(&sideMenu).Error; err != nil {
			return err
		}
	}

	log.Println("サイドメニューデータを挿入しました")
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
			SideMenuID: 1,
			UserID:     2,
			Rating:     5,
			Title:      "とても美味しかった！",
			Comment:    "新鮮な野菜で、ドレッシングも絶品でした。また食べたいです。",
			IsVerified: true,
		},
		{
			SideMenuID: 1,
			UserID:     3,
			Rating:     4,
			Title:      "良いサラダ",
			Comment:    "野菜が新鮮で美味しかったです。",
			IsVerified: true,
		},
		{
			SideMenuID: 2,
			UserID:     2,
			Rating:     5,
			Title:      "最高のポテトフライ",
			Comment:    "カリッとしていて、塩加減も絶妙でした。",
			IsVerified: true,
		},
		{
			SideMenuID: 2,
			UserID:     4,
			Rating:     3,
			Title:      "普通のポテトフライ",
			Comment:    "特に特徴はありませんが、美味しかったです。",
			IsVerified: false,
		},
		{
			SideMenuID: 4,
			UserID:     3,
			Rating:     4,
			Title:      "ジューシーなチキン",
			Comment:    "チキンが柔らかくて美味しかったです。",
			IsVerified: true,
		},
		{
			SideMenuID: 6,
			UserID:     4,
			Rating:     5,
			Title:      "エビがプリプリ",
			Comment:    "エビが新鮮で、衣もサクサクでした。",
			IsVerified: true,
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

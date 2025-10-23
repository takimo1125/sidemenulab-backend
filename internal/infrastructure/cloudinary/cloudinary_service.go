package cloudinary

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService struct {
	cld *cloudinary.Cloudinary
}

type UploadResult struct {
	PublicID  string `json:"public_id"`
	URL       string `json:"url"`
	SecureURL string `json:"secure_url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Format    string `json:"format"`
	Bytes     int    `json:"bytes"`
}

func NewCloudinaryService(cloudName, apiKey, apiSecret string) (*CloudinaryService, error) {
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, fmt.Errorf("Cloudinaryの初期化に失敗しました: %w", err)
	}

	return &CloudinaryService{
		cld: cld,
	}, nil
}

func (s *CloudinaryService) UploadImage(ctx context.Context, file io.Reader, folder string, publicID string) (*UploadResult, error) {
	// アップロードオプションを設定
	uploadOptions := uploader.UploadParams{
		Folder:    folder,
		PublicID:  publicID,
		Overwrite: &[]bool{true}[0],
		ResourceType: "image",
		Transformation: "f_auto,q_auto", // 自動フォーマット変換と品質最適化
	}

	// 画像をアップロード
	result, err := s.cld.Upload.Upload(ctx, file, uploadOptions)
	if err != nil {
		return nil, fmt.Errorf("画像のアップロードに失敗しました: %w", err)
	}

	return &UploadResult{
		PublicID:  result.PublicID,
		URL:       result.URL,
		SecureURL: result.SecureURL,
		Width:     result.Width,
		Height:    result.Height,
		Format:    result.Format,
		Bytes:     result.Bytes,
	}, nil
}

func (s *CloudinaryService) UploadImageFromFile(ctx context.Context, filePath string, folder string, publicID string) (*UploadResult, error) {
	// アップロードオプションを設定
	uploadOptions := uploader.UploadParams{
		Folder:    folder,
		PublicID:  publicID,
		Overwrite: &[]bool{true}[0],
		ResourceType: "image",
		Transformation: "f_auto,q_auto", // 自動フォーマット変換と品質最適化
	}

	// 画像をアップロード
	result, err := s.cld.Upload.Upload(ctx, filePath, uploadOptions)
	if err != nil {
		return nil, fmt.Errorf("画像のアップロードに失敗しました: %w", err)
	}

	return &UploadResult{
		PublicID:  result.PublicID,
		URL:       result.URL,
		SecureURL: result.SecureURL,
		Width:     result.Width,
		Height:    result.Height,
		Format:    result.Format,
		Bytes:     result.Bytes,
	}, nil
}

func (s *CloudinaryService) DeleteImage(ctx context.Context, publicID string) error {
	_, err := s.cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     publicID,
		ResourceType: "image",
	})
	if err != nil {
		return fmt.Errorf("画像の削除に失敗しました: %w", err)
	}
	return nil
}

// GeneratePublicID 一意のPublicIDを生成
func GeneratePublicID(reviewID uint, timestamp int64, index int) string {
	return fmt.Sprintf("review_%d_%d_%d", reviewID, timestamp, index)
}

// GenerateFolderPath フォルダパスを生成
func GenerateFolderPath() string {
	return fmt.Sprintf("sidemenulab/reviews/%d", time.Now().Year())
}

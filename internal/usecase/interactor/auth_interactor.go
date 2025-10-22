package interactor

import (
	"errors"
	"time"

	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/domain/repository"

	"github.com/golang-jwt/jwt/v5"
)

type AuthInteractor struct {
	userRepo repository.UserRepository
	jwtSecret string
}

func NewAuthInteractor(userRepo repository.UserRepository, jwtSecret string) *AuthInteractor {
	return &AuthInteractor{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (a *AuthInteractor) SignUp(req *entity.SignUpRequest) (*entity.AuthResponse, error) {
	// メールアドレスの重複チェック
	existingUser, err := a.userRepo.GetByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("このメールアドレスは既に使用されています")
	}

	// 新しいユーザーを作成
	user := &entity.User{
		Email: req.Email,
		Name:  req.Name,
	}

	// パスワードをハッシュ化
	if err := user.HashPassword(req.Password); err != nil {
		return nil, err
	}

	// ユーザーをデータベースに保存
	if err := a.userRepo.Create(user); err != nil {
		return nil, err
	}

	// JWTトークンを生成
	token, err := a.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &entity.AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

func (a *AuthInteractor) SignIn(req *entity.SignInRequest) (*entity.AuthResponse, error) {
	// ユーザーをメールアドレスで検索
	user, err := a.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("メールアドレスまたはパスワードが正しくありません")
	}

	// パスワードを検証
	if !user.CheckPassword(req.Password) {
		return nil, errors.New("メールアドレスまたはパスワードが正しくありません")
	}

	// JWTトークンを生成
	token, err := a.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &entity.AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

func (a *AuthInteractor) generateToken(user *entity.User) (*entity.AuthToken, error) {
	// アクセストークンの生成
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	accessTokenString, err := accessToken.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return nil, err
	}

	// リフレッシュトークンの生成
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return nil, err
	}

	return &entity.AuthToken{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresAt:    time.Now().Add(24 * time.Hour),
		TokenType:    "Bearer",
	}, nil
}

package handler

import (
	"fmt"
	"net/http"
	"strings"

	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	authUseCase interfaces.AuthUseCase
}

func NewAuthHandler(authUseCase interfaces.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// SignUp ユーザー登録
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req entity.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "リクエストの形式が正しくありません",
			"details": err.Error(),
		})
		return
	}

	response, err := h.authUseCase.SignUp(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ユーザー登録が完了しました",
		"data":    response,
	})
}

// SignIn ユーザーログイン
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req entity.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "リクエストの形式が正しくありません",
			"details": err.Error(),
		})
		return
	}

	response, err := h.authUseCase.SignIn(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ログインに成功しました",
		"data":    response,
	})
}

// DebugToken JWTトークンのデバッグ用エンドポイント
func (h *AuthHandler) DebugToken(c *gin.Context) {
	// Authorizationヘッダーからトークンを取得
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "認証トークンが提供されていません"})
		return
	}

	// "Bearer "プレフィックスを除去
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効な認証ヘッダー形式です"})
		return
	}

	// JWTトークンを解析（署名検証なし）
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "トークンの解析に失敗しました"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.JSON(http.StatusOK, gin.H{
			"claims": claims,
			"user_id_type": fmt.Sprintf("%T", claims["user_id"]),
			"user_id_value": claims["user_id"],
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "クレームの取得に失敗しました"})
	}
}

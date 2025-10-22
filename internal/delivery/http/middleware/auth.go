package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware JWT認証ミドルウェア
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorizationヘッダーからトークンを取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "認証トークンが提供されていません"})
			c.Abort()
			return
		}

		// "Bearer "プレフィックスを除去
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "無効な認証ヘッダー形式です"})
			c.Abort()
			return
		}

		// JWTトークンを解析
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 署名方法を確認
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "無効な認証トークンです"})
			c.Abort()
			return
		}

		// トークンの有効性を確認
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "認証トークンが無効です"})
			c.Abort()
			return
		}

		// クレームからユーザー情報を取得
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID, ok := claims["user_id"].(float64)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDが取得できません"})
				c.Abort()
				return
			}

			email, ok := claims["email"].(string)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "メールアドレスが取得できません"})
				c.Abort()
				return
			}

			// コンテキストにユーザー情報を設定
			c.Set("user_id", uint(userID))
			c.Set("user_email", email)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "認証トークンの解析に失敗しました"})
			c.Abort()
			return
		}

		c.Next()
	}
}

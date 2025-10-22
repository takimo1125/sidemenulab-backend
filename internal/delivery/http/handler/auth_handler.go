package handler

import (
	"net/http"

	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
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

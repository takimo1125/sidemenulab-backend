package handler

import (
	"net/http"
	"strconv"

	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

type SideMenuHandler struct {
	sideMenuUseCase interfaces.SideMenuUseCase
}

func NewSideMenuHandler(sideMenuUseCase interfaces.SideMenuUseCase) *SideMenuHandler {
	return &SideMenuHandler{
		sideMenuUseCase: sideMenuUseCase,
	}
}

// CreateStore 店舗作成
func (h *SideMenuHandler) CreateStore(c *gin.Context) {
	var req entity.CreateStoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store, err := h.sideMenuUseCase.CreateStore(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "店舗が作成されました", "data": store})
}

// GetStoreByID 店舗詳細取得
func (h *SideMenuHandler) GetStoreByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	store, err := h.sideMenuUseCase.GetStoreByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": store})
}

// GetAllStores 店舗一覧取得
func (h *SideMenuHandler) GetAllStores(c *gin.Context) {
	stores, err := h.sideMenuUseCase.GetAllStores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": stores})
}

// CreateSideMenu サイドメニュー作成
func (h *SideMenuHandler) CreateSideMenu(c *gin.Context) {
	var req entity.CreateSideMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sideMenu, err := h.sideMenuUseCase.CreateSideMenu(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "サイドメニューが作成されました", "data": sideMenu})
}

// GetSideMenuByID サイドメニュー詳細取得
func (h *SideMenuHandler) GetSideMenuByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	sideMenu, err := h.sideMenuUseCase.GetSideMenuByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sideMenu})
}

// GetSideMenusByStoreID 店舗別サイドメニュー一覧取得
func (h *SideMenuHandler) GetSideMenusByStoreID(c *gin.Context) {
	storeIDStr := c.Param("storeId")
	storeID, err := strconv.ParseUint(storeIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効な店舗IDです"})
		return
	}

	sideMenus, err := h.sideMenuUseCase.GetSideMenusByStoreID(uint(storeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sideMenus})
}

// GetAllSideMenus 全サイドメニュー一覧取得
func (h *SideMenuHandler) GetAllSideMenus(c *gin.Context) {
	sideMenus, err := h.sideMenuUseCase.GetAllSideMenus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sideMenus})
}

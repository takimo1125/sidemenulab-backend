package interactor

import (
	"fmt"

	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/domain/repository"
	"sidemenulab-backend/internal/usecase/interfaces"
)

type SideMenuInteractor struct {
	storeRepo    repository.StoreRepository
	sideMenuRepo repository.SideMenuRepository
}

func NewSideMenuInteractor(storeRepo repository.StoreRepository, sideMenuRepo repository.SideMenuRepository) interfaces.SideMenuUseCase {
	return &SideMenuInteractor{
		storeRepo:    storeRepo,
		sideMenuRepo: sideMenuRepo,
	}
}

func (i *SideMenuInteractor) CreateStore(req *entity.CreateStoreRequest) (*entity.Store, error) {
	store := &entity.Store{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}

	if err := i.storeRepo.CreateStore(store); err != nil {
		return nil, fmt.Errorf("店舗の作成に失敗しました: %w", err)
	}

	return store, nil
}

func (i *SideMenuInteractor) GetStoreByID(id uint) (*entity.Store, error) {
	store, err := i.storeRepo.GetStoreByID(id)
	if err != nil {
		return nil, fmt.Errorf("店舗の取得に失敗しました: %w", err)
	}
	return store, nil
}

func (i *SideMenuInteractor) GetAllStores() ([]*entity.Store, error) {
	stores, err := i.storeRepo.GetAllStores()
	if err != nil {
		return nil, fmt.Errorf("店舗一覧の取得に失敗しました: %w", err)
	}
	return stores, nil
}

func (i *SideMenuInteractor) CreateSideMenu(req *entity.CreateSideMenuRequest) (*entity.SideMenu, error) {
	// 店舗の存在確認
	_, err := i.storeRepo.GetStoreByID(req.StoreID)
	if err != nil {
		return nil, fmt.Errorf("指定された店舗が見つかりません: %w", err)
	}

	sideMenu := &entity.SideMenu{
		StoreID:     req.StoreID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	if err := i.sideMenuRepo.CreateSideMenu(sideMenu); err != nil {
		return nil, fmt.Errorf("サイドメニューの作成に失敗しました: %w", err)
	}

	// 作成されたサイドメニューを店舗情報と一緒に取得
	createdSideMenu, err := i.sideMenuRepo.GetSideMenuByID(sideMenu.ID)
	if err != nil {
		return nil, fmt.Errorf("作成されたサイドメニューの取得に失敗しました: %w", err)
	}

	return createdSideMenu, nil
}

func (i *SideMenuInteractor) GetSideMenuByID(id uint) (*entity.SideMenu, error) {
	sideMenu, err := i.sideMenuRepo.GetSideMenuByID(id)
	if err != nil {
		return nil, fmt.Errorf("サイドメニューの取得に失敗しました: %w", err)
	}
	return sideMenu, nil
}

func (i *SideMenuInteractor) GetSideMenusByStoreID(storeID uint) ([]*entity.SideMenu, error) {
	sideMenus, err := i.sideMenuRepo.GetSideMenusByStoreID(storeID)
	if err != nil {
		return nil, fmt.Errorf("店舗のサイドメニュー一覧の取得に失敗しました: %w", err)
	}
	return sideMenus, nil
}

func (i *SideMenuInteractor) GetAllSideMenus() ([]*entity.SideMenu, error) {
	sideMenus, err := i.sideMenuRepo.GetAllSideMenus()
	if err != nil {
		return nil, fmt.Errorf("サイドメニュー一覧の取得に失敗しました: %w", err)
	}
	return sideMenus, nil
}

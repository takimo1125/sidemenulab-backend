package interfaces

import "sidemenulab-backend/internal/domain/entity"

type SideMenuUseCase interface {
	CreateStore(req *entity.CreateStoreRequest) (*entity.Store, error)
	GetStoreByID(id uint) (*entity.Store, error)
	GetAllStores() ([]*entity.Store, error)
	CreateSideMenu(req *entity.CreateSideMenuRequest) (*entity.SideMenu, error)
	GetSideMenuByID(id uint) (*entity.SideMenu, error)
	GetSideMenusByStoreID(storeID uint) ([]*entity.SideMenu, error)
	GetAllSideMenus() ([]*entity.SideMenu, error)
}

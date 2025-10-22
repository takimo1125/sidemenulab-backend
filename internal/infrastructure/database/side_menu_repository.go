package database

import (
	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/domain/repository"

	"gorm.io/gorm"
)

type SideMenuRepository struct {
	db *gorm.DB
}

func NewSideMenuRepository(db *gorm.DB) repository.SideMenuRepository {
	return &SideMenuRepository{db: db}
}

func (r *SideMenuRepository) CreateSideMenu(sideMenu *entity.SideMenu) error {
	return r.db.Create(sideMenu).Error
}

func (r *SideMenuRepository) GetSideMenuByID(id uint) (*entity.SideMenu, error) {
	var sideMenu entity.SideMenu
	if err := r.db.Preload("Store").First(&sideMenu, id).Error; err != nil {
		return nil, err
	}
	return &sideMenu, nil
}

func (r *SideMenuRepository) GetSideMenusByStoreID(storeID uint) ([]*entity.SideMenu, error) {
	var sideMenus []*entity.SideMenu
	if err := r.db.Preload("Store").Where("store_id = ?", storeID).Find(&sideMenus).Error; err != nil {
		return nil, err
	}
	return sideMenus, nil
}

func (r *SideMenuRepository) GetAllSideMenus() ([]*entity.SideMenu, error) {
	var sideMenus []*entity.SideMenu
	if err := r.db.Preload("Store").Find(&sideMenus).Error; err != nil {
		return nil, err
	}
	return sideMenus, nil
}

func (r *SideMenuRepository) UpdateSideMenu(sideMenu *entity.SideMenu) error {
	return r.db.Save(sideMenu).Error
}

func (r *SideMenuRepository) DeleteSideMenu(id uint) error {
	return r.db.Delete(&entity.SideMenu{}, id).Error
}

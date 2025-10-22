package database

import (
	"sidemenulab-backend/internal/domain/entity"
	"sidemenulab-backend/internal/domain/repository"

	"gorm.io/gorm"
)

type StoreRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) repository.StoreRepository {
	return &StoreRepository{db: db}
}

func (r *StoreRepository) CreateStore(store *entity.Store) error {
	return r.db.Create(store).Error
}

func (r *StoreRepository) GetStoreByID(id uint) (*entity.Store, error) {
	var store entity.Store
	if err := r.db.First(&store, id).Error; err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *StoreRepository) GetAllStores() ([]*entity.Store, error) {
	var stores []*entity.Store
	if err := r.db.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

func (r *StoreRepository) UpdateStore(store *entity.Store) error {
	return r.db.Save(store).Error
}

func (r *StoreRepository) DeleteStore(id uint) error {
	return r.db.Delete(&entity.Store{}, id).Error
}

package repository

import (
	"gorm.io/gorm"
	"next_device/backend/models"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db}
}

func (ar *AddressRepository) CreateAddress(address *models.Address) error {
	if result := ar.db.Create(address); result.Error != nil {
		return result.Error
	}
	return nil
}

func (ar *AddressRepository) GetAddressByUserId(id uint) (*models.Address, error) {
	var address *models.Address
	if result := ar.db.First(&address, "user_id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return address, nil
}

func (ar *AddressRepository) UpdateAddress(address *models.Address) error {
	if result := ar.db.Save(address); result.Error != nil {
		return result.Error
	}
	return nil
}

func (ar *AddressRepository) DeleteAddressByUserId(id uint) error {
	if result := ar.db.Delete(&models.Address{}, "user_id = ?", id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (ar *AddressRepository) DeleteAddressById(id uint) error {
	if result := ar.db.Delete(&models.Address{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

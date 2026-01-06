package repository

import (
	"github.com/MatheusMikio/schemas"
	"gorm.io/gorm"
)

type IClientRepository interface {
	GetAll() ([]schemas.Client, error)
	// GetClientById(id uint) (schemas.Client, error)
}

type ClientRepository struct {
	Db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{Db: db}
}

func (repository *ClientRepository) GetAll() ([]schemas.Client, error) {
	clients := []schemas.Client{}
	if err := repository.Db.Preload("Appointments").Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

package repository

import (
	"github.com/MatheusMikio/schemas"
	"gorm.io/gorm"
)

type IClientRepository interface {
	GetAll() (*[]schemas.Client, error)
	GetById(id uint) (*schemas.Client, error)
}

type ClientRepository struct {
	Db *gorm.DB
}

func NewClientRepository(db *gorm.DB) IClientRepository {
	return &ClientRepository{
		Db: db,
	}
}

func (repository *ClientRepository) GetAll() (*[]schemas.Client, error) {
	clients := []schemas.Client{}
	if err := repository.Db.Preload("Appointments").Find(&clients).Error; err != nil {
		return nil, err
	}
	return &clients, nil
}

func (repository *ClientRepository) GetById(id uint) (*schemas.Client, error) {
	client := schemas.Client{}
	if err := repository.Db.Preload("Appointments").First(&client, id).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

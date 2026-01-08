package repository

import (
	"github.com/MatheusMikio/schemas"
	"gorm.io/gorm"
)

type IClientRepository interface {
	GetAll() (*[]schemas.Client, error)
	GetById(id uint) (*schemas.Client, error)
	GetDb() *gorm.DB
	Create(client *schemas.Client) error
}

type ClientRepository struct {
	Db *gorm.DB
}

func NewClientRepository(db *gorm.DB) IClientRepository {
	return &ClientRepository{
		Db: db,
	}
}

func (r *ClientRepository) GetDb() *gorm.DB {
	return r.Db
}

func (r *ClientRepository) GetAll() (*[]schemas.Client, error) {
	clients := []schemas.Client{}
	if err := r.Db.Preload("Appointments").Find(&clients).Error; err != nil {
		return nil, err
	}
	return &clients, nil
}

func (r *ClientRepository) GetById(id uint) (*schemas.Client, error) {
	client := schemas.Client{}
	if err := r.Db.Preload("Appointments").First(&client, id).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *ClientRepository) Create(client *schemas.Client) error {
	if err := r.Db.Create(client).Error; err != nil {
		return err
	}
	return nil
}

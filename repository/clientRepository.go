package repository

import (
	"github.com/MatheusMikio/schemas"
	"gorm.io/gorm"
)

type IClientRepository interface {
	GetAll() (*[]schemas.Client, error)
	GetById(id uint) (*schemas.Client, error)
	GetByEmail(email string) (*schemas.Client, error)
	GetByGoogleId(googleId string) (*schemas.Client, error)
	GetDb() *gorm.DB
	Create(client *schemas.Client) error
	Update(client *schemas.Client) error
	Delete(client *schemas.Client) error
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
	if err := r.Db.Preload("Appointments").Preload("CardData").First(&client, id).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *ClientRepository) GetByEmail(email string) (*schemas.Client, error) {
	client := schemas.Client{}
	if err := r.Db.Preload("Appointments").Where("email = ?", email).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *ClientRepository) GetByGoogleId(googleId string) (*schemas.Client, error) {
	client := schemas.Client{}
	if err := r.Db.Preload("Appointments").Where("google_id = ?", googleId).First(&client).Error; err != nil {
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

func (r *ClientRepository) Update(client *schemas.Client) error {
	if err := r.Db.Session(&gorm.Session{FullSaveAssociations: true}).Save(client).Error; err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) Delete(client *schemas.Client) error {
	if err := r.Db.Delete(client, client.ID).Error; err != nil {
		return err
	}
	return nil
}

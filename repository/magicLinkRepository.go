package repository

import (
	"time"

	"github.com/MatheusMikio/schemas"
	"gorm.io/gorm"
)

type IMagicLinkRepository interface {
	Create(m *schemas.MagicLink) error
	GetByToken(token string) (*schemas.MagicLink, error)
	MarkUsed(id uint) error
}

type MagicLinkRepository struct {
	Db *gorm.DB
}

func NewMagicLinkRepository(db *gorm.DB) IMagicLinkRepository {
	return &MagicLinkRepository{
		Db: db,
	}
}

func (r *MagicLinkRepository) Create(m *schemas.MagicLink) error {
	if err := r.Db.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (r *MagicLinkRepository) GetByToken(token string) (*schemas.MagicLink, error) {
	magicLink := schemas.MagicLink{}
	if err := r.Db.Where("token = ? and used = ? and expires_at > ?", token, false, time.Now()).First(&magicLink).Error; err != nil {
		return nil, err
	}
	return &magicLink, nil
}

func (r *MagicLinkRepository) MarkUsed(id uint) error {
	if err := r.Db.Model(&schemas.MagicLink{}).Where("id = ?", id).Update("used", true).Error; err != nil {
		return err
	}
	return nil
}

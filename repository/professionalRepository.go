package repository

import (
	"github.com/MatheusMikio/schemas"
	"gorm.io/gorm"
)

type IProfessionalRepository interface {
	GetByEmail(email string) (*schemas.Professional, error)
	GetByGoogleId(googleId string) (*schemas.Professional, error)
	Update(professional *schemas.Professional) error
}

type ProfessionalRepository struct {
	Db *gorm.DB
}

func NewProfessionalRepository(db *gorm.DB) IProfessionalRepository {
	return &ProfessionalRepository{
		Db: db,
	}
}

func (p *ProfessionalRepository) GetByEmail(email string) (*schemas.Professional, error) {
	professional := schemas.Professional{}
	if err := p.Db.Preload("Appointments").Where("email = ?", email).First(&professional).Error; err != nil {
		return nil, err
	}
	return &professional, nil
}

func (p *ProfessionalRepository) GetByGoogleId(googleId string) (*schemas.Professional, error) {
	professional := schemas.Professional{}
	if err := p.Db.Preload("Appointments").Where("google_id = ?", googleId).First(&professional).Error; err != nil {
		return nil, err
	}
	return &professional, nil
}

func (r *ProfessionalRepository) Update(professional *schemas.Professional) error {
	if err := r.Db.Session(&gorm.Session{FullSaveAssociations: true}).Save(professional).Error; err != nil {
		return err
	}
	return nil
}

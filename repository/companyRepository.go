package repository

import (
	"github.com/MatheusMikio/schemas"
	"gorm.io/gorm"
)

type ICompanyRespository interface {
	GetByEmail(email string) (*schemas.Company, error)
	Update(c *schemas.Company) (*schemas.Company, error)
}

type CompanyRepository struct {
	Db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) ICompanyRespository {
	return &CompanyRepository{
		Db: db,
	}
}

func (c *CompanyRepository) GetByEmail(email string) (*schemas.Company, error) {
	panic("unimplemented")
}

func (*CompanyRepository) Update(c *schemas.Company) (*schemas.Company, error) {
	panic("unimplemented")
}

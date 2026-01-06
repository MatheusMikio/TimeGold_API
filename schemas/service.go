package schemas

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name         string          `json:"name"`
	Duration     string          `json:"duration"`
	Price        decimal.Decimal `json:"price" gorm:"type:decimal(10,2)"`
	Description  string          `json:"description"`
	CompanyID    uint            `json:"companyId"`
	Company      Company         `json:"-" gorm:"foreignKey:CompanyID"`
	Appointments []Scheduling    `json:"appointments" gorm:"foreignKey:ServiceID"`
}

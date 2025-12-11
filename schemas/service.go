package schemas

import "github.com/shopspring/decimal"

type Service struct {
	Name      string          `json:"name"`
	Duration  string          `json:"duration"`
	Price     decimal.Decimal `json:"price"`
	CompanyID uint            `json:"companyId"`
	Company   Company         `json:"-" gorm:"foreignKey:CompanyID"`
}

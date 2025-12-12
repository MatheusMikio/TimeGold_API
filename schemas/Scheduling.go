package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Scheduling struct {
	gorm.Model
	Date           time.Time    `json:"date"`
	ServiceID      uint         `json:"serviceId"`
	Service        Service      `json:"service" gorm:"foreignKey:ServiceID"`
	ProfessionalID uint         `json:"professionalId"`
	Professional   Professional `json:"professional" gorm:"foreignKey:ProfessionalID"`
	CompanyID      uint         `json:"companyId"`
	Company        Company      `json:"company" gorm:"foreignKey:CompanyID"`
	ClientID       uint         `json:"clientId"`
	Client         Client       `json:"client" gorm:"foreignKey:ClientID"`
}

package schemas

import (
	"github.com/MatheusMikio/models/base"
)

type Professional struct {
	base.BaseUser
	CompanyID    uint         `json:"companyId"`
	Company      Company      `json:"-" gorm:"foreignKey:CompanyID"`
	Appointments []Scheduling `json:"appointments"`
}

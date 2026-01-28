package schemas

import (
	"github.com/MatheusMikio/models/base"
)

type ProfessionalRole string

const (
	RoleProfessional ProfessionalRole = "PROFESSIONAL"
	RoleAdmin        ProfessionalRole = "ADMIN"
)

type Professional struct {
	base.BaseUser
	CompanyID    uint             `json:"companyId"`
	Company      Company          `json:"-" gorm:"foreignKey:CompanyID"`
	Role         ProfessionalRole `json:"role" gorm:"type:varchar(20);default:'PROFESSIONAL'"`
	Appointments []Scheduling     `json:"appointments"`
}

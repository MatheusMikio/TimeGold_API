package professional

import (
	"github.com/MatheusMikio/schemas"
)

type ProfessionalResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Cpf       string `json:"cpf"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CompanyID uint   `json:"companyId"`
	Role      schemas.ProfessionalRole
}

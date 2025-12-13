package base

import (
	"gorm.io/gorm"
)

type BaseUser struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Cpf       string `json:"cpf"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

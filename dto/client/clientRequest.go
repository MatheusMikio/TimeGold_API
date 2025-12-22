package client

import (
	"github.com/MatheusMikio/models"
	"github.com/MatheusMikio/schemas"
	"github.com/MatheusMikio/utils"
	"gorm.io/gorm"
)

type ClientRequest struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Cpf          string `json:"cpf"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	StripeCardId string `json:"stripeCardId"`
	CardBrand    string `json:"cardBrand"`
	CardLast4    string `json:"cardLast4"`
	CardExpMonth int    `json:"cardExpMonth"`
	CardExpYear  int    `json:"cardExpYear"`
}

type UpdateClientRequest struct {
	Id           int64  `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Cpf          string `json:"cpf"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	CardBrand    string `json:"cardBrand"`
	CardLast4    string `json:"cardLast4"`
	CardExpMonth int    `json:"cardExpMonth"`
	CardExpYear  int    `json:"cardExpYear"`
}

func (r *ClientRequest) Validate(db *gorm.DB) []*models.ErrorMessage {
	errorMessages := []*models.ErrorMessage{}

	if len(r.FirstName) < 3 || len(r.FirstName) > 20 {
		errorMessages = append(errorMessages, models.CreateErrorMessage("FirstName", "The first name must have a minimum of 3 characters and a maximum of 20."))
	}

	if len(r.LastName) < 3 || len(r.LastName) > 20 {
		errorMessages = append(errorMessages, models.CreateErrorMessage("LastName", "The last name must have a minimum of 3 characters and a maximum of 20."))
	}

	errorMessages = append(errorMessages, utils.ValidateEmail[schemas.Client](r.Email, db)...)

	errorMessages = append(errorMessages, utils.ValidateCpf[schemas.Client](r.Cpf, db)...)

	return errorMessages
}

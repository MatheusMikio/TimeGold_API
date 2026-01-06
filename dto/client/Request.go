package client

import (
	"strings"
	"time"

	"github.com/MatheusMikio/models"
	"github.com/MatheusMikio/schemas"
	"github.com/MatheusMikio/utils"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentmethod"
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
	StripeCardId string `json:"stripeCardId"`
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

	errorMessages = append(errorMessages, utils.ValidatePhone[schemas.Client](r.Phone, db)...)

	return errorMessages
}

func (r *ClientRequest) ValidateCardStripe(stripeKey string) []*models.ErrorMessage {
	errorsMessage := []*models.ErrorMessage{}

	if r.StripeCardId == "" {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("Card", "StripeCardId is required"))
		return errorsMessage
	}

	if len(r.CardLast4) != 4 {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("CardLast4", "CardLast4 must be 4 digits"))
	}

	if r.CardExpMonth < 1 || r.CardExpMonth > 12 {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("CardExpMonth", "CardExpMonth must be between 1 and 12"))
	}

	now := time.Now()
	if r.CardExpYear < now.Year() {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("CardExpYear", "CardExpYear must be this year or later"))

	}

	if len(errorsMessage) > 0 {
		return errorsMessage
	}

	stripe.Key = stripeKey
	pm, err := paymentmethod.Get(r.StripeCardId, nil)
	if err != nil {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("Stripe", "Unable to fetch payment method from Stripe: "+err.Error()))
		return errorsMessage
	}

	if pm.Card == nil {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("Stripe", "Fetched object has no card details"))
		return errorsMessage
	}

	if !strings.EqualFold(string(pm.Card.Brand), r.CardBrand) {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("CardBrand", "CardBrand does not match Stripe record"))
	}
	if pm.Card.Last4 != r.CardLast4 {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("CardLast4", "CardLast4 does not match Stripe record"))
	}
	if int(pm.Card.ExpMonth) != r.CardExpMonth || int(pm.Card.ExpYear) != r.CardExpYear {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("CardExp", "Expiry month/year do not match Stripe record"))
	}

	return errorsMessage
}

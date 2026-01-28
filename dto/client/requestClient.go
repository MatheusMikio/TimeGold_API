package client

import (
	"github.com/MatheusMikio/models"
	"github.com/MatheusMikio/schemas"
	"github.com/MatheusMikio/utils"
	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/paymentmethod"
	"gorm.io/gorm"
)

type ClientRequest struct {
	FirstName             string `json:"firstName"`
	LastName              string `json:"lastName"`
	Cpf                   string `json:"cpf"`
	Email                 string `json:"email"`
	Phone                 string `json:"phone"`
	StripePaymentMethodId string `json:"stripePaymentMethodId"`
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

func (r *ClientRequest) ValidateAndFetchCard(stripeKey string) (*models.CardData, []*models.ErrorMessage) {
	errorsMessage := []*models.ErrorMessage{}

	if r.StripePaymentMethodId == "" {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("Card", "Payment method ID is required"))
		return nil, errorsMessage
	}

	stripe.Key = stripeKey

	pm, err := paymentmethod.Get(r.StripePaymentMethodId, nil)
	if err != nil {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("Stripe", "Invalid payment method:"+err.Error()))
		return nil, errorsMessage
	}

	cardData := &models.CardData{
		StripeCardId: pm.ID,
		CardBrand:    string(pm.Card.Brand),
		CardLast4:    pm.Card.Last4,
		CardExpMonth: int(pm.Card.ExpMonth),
		CardExpYear:  int(pm.Card.ExpYear),
	}
	return cardData, nil
}

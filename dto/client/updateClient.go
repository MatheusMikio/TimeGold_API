package client

import (
	"errors"

	"github.com/MatheusMikio/models"
	"github.com/MatheusMikio/schemas"
	"github.com/MatheusMikio/utils"
	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/paymentmethod"
	"gorm.io/gorm"
)

type UpdateClientRequest struct {
	Id                    uint   `json:"id"`
	FirstName             string `json:"firstName"`
	LastName              string `json:"lastName"`
	Cpf                   string `json:"cpf"`
	Email                 string `json:"email"`
	Phone                 string `json:"phone"`
	StripePaymentMethodId string `json:"stripePaymentMethodId"` // ← APENAS O ID (opcional no update)
}

func (ucr *UpdateClientRequest) Validate(db *gorm.DB) []*models.ErrorMessage {
	errorMessages := []*models.ErrorMessage{}

	clientDb := schemas.Client{}
	if err := db.First(&clientDb, ucr.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMessages = append(errorMessages, models.CreateErrorMessage("Client", "Not found"))
			return errorMessages
		}
		errorMessages = append(errorMessages, models.CreateErrorMessage("System", err.Error()))
		return errorMessages
	}

	if len(ucr.FirstName) < 3 || len(ucr.FirstName) > 20 {
		errorMessages = append(errorMessages, models.CreateErrorMessage("FirstName", "The first name must have a minimum of 3 characters and a maximum of 20."))
	}

	if len(ucr.LastName) < 3 || len(ucr.LastName) > 20 {
		errorMessages = append(errorMessages, models.CreateErrorMessage("LastName", "The last name must have a minimum of 3 characters and a maximum of 20."))
	}

	errorMessages = append(errorMessages, utils.ValidateEmailUpdate[schemas.Client](ucr.Email, ucr.Id, db)...)
	errorMessages = append(errorMessages, utils.ValidateCpfUpdate[schemas.Client](ucr.Cpf, ucr.Id, db)...)
	errorMessages = append(errorMessages, utils.ValidatePhoneUpdate[schemas.Client](ucr.Phone, ucr.Id, db)...)

	return errorMessages
}

func (ucr *UpdateClientRequest) HasCardChanged(client *schemas.Client) bool {
	if ucr.StripePaymentMethodId == "" {
		return false
	}

	if client.CardData == nil {
		return true
	}

	return ucr.StripePaymentMethodId != client.CardData.StripeCardId
}

func (ucr *UpdateClientRequest) ValidateAndFetchCard(stripeKey string) (*models.CardData, []*models.ErrorMessage) {
	errorsMessage := []*models.ErrorMessage{}

	if ucr.StripePaymentMethodId == "" {
		return nil, nil
	}

	stripe.Key = stripeKey

	pm, err := paymentmethod.Get(ucr.StripePaymentMethodId, nil)
	if err != nil {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("Stripe", "Invalid payment method: "+err.Error()))
		return nil, errorsMessage
	}

	if pm.Card == nil {
		errorsMessage = append(errorsMessage, models.CreateErrorMessage("Stripe", "Payment method is not a card"))
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

func (ucr *UpdateClientRequest) MergeInto(client *schemas.Client) {
	client.FirstName = ucr.FirstName
	client.LastName = ucr.LastName
	client.Cpf = ucr.Cpf
	client.Email = ucr.Email
	client.Phone = ucr.Phone
}

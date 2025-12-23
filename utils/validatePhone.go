package utils

import (
	"github.com/MatheusMikio/models"
	validate "github.com/nyaruka/phonenumbers"
	"gorm.io/gorm"
)

func ValidatePhone[T any](phone string, db *gorm.DB) []*models.ErrorMessage {
	errorMessages := []*models.ErrorMessage{}

	if phone == "" {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Phone", "Phone cannot be empty"))
		return errorMessages
	}

	num, err := validate.Parse(phone, "BR")

	if err != nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Phone", "Invalid phone"))
		return errorMessages
	}

	if !validate.IsValidNumber(num) {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Phone", "Invalid phone"))
	}

	var existing T
	if err := db.Where("phone = ?", phone).First(existing).Error; err != nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Phone", "Invalid phone"))
	}
	return errorMessages
}

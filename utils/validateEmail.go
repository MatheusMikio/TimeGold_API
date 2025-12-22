package utils

import (
	"github.com/MatheusMikio/models"
	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

func ValidateEmail[T any](email string, db *gorm.DB) []*models.ErrorMessage {
	errorMessages := []*models.ErrorMessage{}

	if email == "" {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Email", "The email address cannot be empty."))
		return errorMessages
	}

	if err := checkmail.ValidateFormat(email); err != nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Email", "Invalid format"))
		return errorMessages
	}

	if err := checkmail.ValidateHost(email); err != nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Email", "Email unavailable."))
		return errorMessages
	}

	var existing T
	if err := db.Where("email = ?", email).First(&existing).Error; err == nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Email", "Email unavailable."))
	}

	return errorMessages
}

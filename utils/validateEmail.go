package utils

import (
	"errors"

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
	err := db.Where("email = ?", email).First(&existing).Error
	if err == nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Email", "Email unavailable."))
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Db", err.Error()))
	}

	return errorMessages
}

func ValidateEmailUpdate[T any](email string, id uint, db *gorm.DB) []*models.ErrorMessage {
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
	err := db.Where("email = ? and id != ?", email, id).First(&existing).Error
	if err == nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Email", "Email unavailable."))
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Db", err.Error()))
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Db", err.Error()))
	}

	return errorMessages
}

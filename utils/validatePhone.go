package utils

import (
	"errors"

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
	err = db.Where("phone = ?", phone).First(&existing).Error
	if err == nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Phone", "Phone unavailable."))
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Db", err.Error()))
	}
	return errorMessages
}

func ValidatePhoneUpdate[T any](phone string, id uint, db *gorm.DB) []*models.ErrorMessage {
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
	err = db.Where("phone = ? and id != ?", phone, id).First(&existing).Error
	if err == nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Phone", "Phone already registered"))
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Db", err.Error()))
	}
	return errorMessages
}

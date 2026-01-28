package utils

import (
	"errors"

	"github.com/MatheusMikio/models"
	cpfPkg "github.com/jfelipearaujo/cpfcnpj/cpf"

	"gorm.io/gorm"
)

func ValidateCpf[T any](cpf string, db *gorm.DB) []*models.ErrorMessage {
	errorMessages := []*models.ErrorMessage{}
	svc := cpfPkg.New()

	if cpf == "" {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Cpf", "CPF cannot be empty"))
		return errorMessages
	}

	if err := svc.IsValid(cpf); err != nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Cpf", "Invalid CPF"))
	}

	var existing T
	err := db.Where("cpf = ?", cpf).First(&existing).Error
	if err == nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Cpf", "Cpf unavailable."))
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Db", err.Error()))
	}

	return errorMessages
}

func ValidateCpfUpdate[T any](cpf string, id uint, db *gorm.DB) []*models.ErrorMessage {
	errorMessages := []*models.ErrorMessage{}
	svc := cpfPkg.New()

	if cpf == "" {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Cpf", "CPF cannot be empty"))
		return errorMessages
	}

	if err := svc.IsValid(cpf); err != nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Cpf", "Invalid CPF"))
	}

	var existing T

	err := db.Where("cpf = ? and id != ?", cpf, id).First(&existing).Error
	if err == nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Cpf", "Cpf unavailable."))
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Db", err.Error()))
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Db", err.Error()))
	}

	return errorMessages
}

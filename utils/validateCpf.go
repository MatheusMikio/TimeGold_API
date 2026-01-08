package utils

import (
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
		errorMessages = append(errorMessages, models.CreateErrorMessage("cpf", "Invalid CPF"))
	}

	var existing T
	if err := db.Where("cpf = ?", cpf).First(&existing).Error; err == nil {
		errorMessages = append(errorMessages, models.CreateErrorMessage("Cpf", "Cpf unavailable."))
	}

	return errorMessages
}

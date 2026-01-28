package auth

import "github.com/MatheusMikio/models"

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (r *LoginRequest) Validate() *models.ErrorMessage {
	if r.Email == "" {
		return models.CreateErrorMessage("Email", "Email is required")
	}
	return nil
}

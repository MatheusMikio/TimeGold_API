package client

import "time"

type ClientResponse struct {
	ID                  uint                 `json:"id"`
	FirstName           string               `json:"firstName"`
	LastName            string               `json:"lastName"`
	Cpf                 string               `json:"cpf"`
	Email               string               `json:"email"`
	Phone               string               `json:"phone"`
	CreatedAt           time.Time            `json:"createdAt"`
	UpdatedAt           time.Time            `json:"updatedAt"`
	DeletedAt           time.Time            `json:"deletedAt,omitempty"`
	AppointmentsSummary []AppointmentSummary `json:"appointmentsSummary,omitempty"`
}

type AppointmentSummary struct {
	ID   uint      `json:"id"`
	Date time.Time `json:"date"`
}

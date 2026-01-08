package services

import (
	"github.com/MatheusMikio/config"
	"github.com/MatheusMikio/dto/client"
	"github.com/MatheusMikio/models"
	"github.com/MatheusMikio/models/base"
	"github.com/MatheusMikio/repository"
	"github.com/MatheusMikio/schemas"
)

type IClientService interface {
	GetAll() (*[]client.ClientResponse, error)
	GetById(id uint) (*client.ClientResponse, error)
	Create(clientRequest *client.ClientRequest) []*models.ErrorMessage
}

type ClientService struct {
	Repository repository.IClientRepository
}

func NewClientService(repo repository.IClientRepository) IClientService {
	return &ClientService{
		Repository: repo,
	}
}

func (service *ClientService) GetAll() (*[]client.ClientResponse, error) {
	clients, err := service.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	clientsResponse := make([]client.ClientResponse, 0, len(*clients))
	for _, cli := range *clients {
		appointmentsSummary := make([]client.AppointmentSummary, 0, len(cli.Appointments))

		for _, scheduling := range cli.Appointments {
			appointmentsSummary = append(appointmentsSummary, client.AppointmentSummary{
				ID:   scheduling.ID,
				Date: scheduling.Date,
			})
		}

		clientResponse := client.ClientResponse{
			ID:                  cli.BaseUser.ID,
			FirstName:           cli.FirstName,
			LastName:            cli.LastName,
			Cpf:                 cli.Cpf,
			Email:               cli.Email,
			Phone:               cli.Phone,
			AppointmentsSummary: appointmentsSummary,
		}
		clientsResponse = append(clientsResponse, clientResponse)
	}
	return &clientsResponse, nil
}

func (service *ClientService) GetById(id uint) (*client.ClientResponse, error) {
	cli, err := service.Repository.GetById(id)

	if err != nil {
		return nil, err
	}

	appointmentsSummary := make([]client.AppointmentSummary, 0, len(cli.Appointments))
	for _, scheduling := range cli.Appointments {
		appointmentsSummary = append(appointmentsSummary, client.AppointmentSummary{
			ID:   scheduling.ID,
			Date: scheduling.Date,
		})
	}
	clientReponse := client.ClientResponse{
		ID:                  cli.BaseUser.ID,
		FirstName:           cli.FirstName,
		LastName:            cli.LastName,
		Cpf:                 cli.Cpf,
		Email:               cli.Email,
		Phone:               cli.Phone,
		AppointmentsSummary: appointmentsSummary,
	}
	return &clientReponse, nil
}

func (service *ClientService) Create(clientRequest *client.ClientRequest) []*models.ErrorMessage {
	logger := config.GetLogger("Create (CLIENT)")
	if errorMessage := clientRequest.Validate(service.Repository.GetDb()); len(errorMessage) > 0 {
		logger.Errorf("Validation failed: %d errors found", len(errorMessage))
		return errorMessage
	}

	newClient := &schemas.Client{
		BaseUser: base.BaseUser{
			FirstName: clientRequest.FirstName,
			LastName:  clientRequest.LastName,
			Cpf:       clientRequest.Cpf,
			Email:     clientRequest.Email,
			Phone:     clientRequest.Phone,
		},
		CardData: &models.CardData{
			StripeCardId: clientRequest.StripeCardId,
			CardBrand:    clientRequest.CardBrand,
			CardLast4:    clientRequest.CardLast4,
			CardExpMonth: clientRequest.CardExpMonth,
			CardExpYear:  clientRequest.CardExpYear,
		},
	}

	if err := service.Repository.Create(newClient); err != nil {
		logger.Errorf("Failed to create client in database: %v", err)
		return []*models.ErrorMessage{
			models.CreateErrorMessage("Database", "Failed to create client: "+err.Error()),
		}
	}

	return nil
}

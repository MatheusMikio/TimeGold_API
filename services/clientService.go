package services

import (
	"errors"
	"fmt"

	"github.com/MatheusMikio/config"
	"github.com/MatheusMikio/dto/client"
	"github.com/MatheusMikio/models"
	"github.com/MatheusMikio/models/base"
	"github.com/MatheusMikio/repository"
	"github.com/MatheusMikio/schemas"
	"gorm.io/gorm"
)

type IClientService interface {
	GetAll() (*[]client.ClientResponse, error)
	GetById(id uint) (*client.ClientResponse, error)
	Create(clientRequest *client.ClientRequest) []*models.ErrorMessage
	Update(clientRequest *client.UpdateClientRequest) []*models.ErrorMessage
	Delete(id uint) *models.ErrorMessage
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
	logger := config.GetLogger("Create (CLIENT) service")

	if errorMessage := clientRequest.Validate(service.Repository.GetDb()); len(errorMessage) > 0 {
		logger.Errorf("Validation failed: %d errors found", len(errorMessage))
		return errorMessage
	}

	cardData, err := clientRequest.ValidateAndFetchCard(config.GetStripeKey())
	if len(err) > 0 {
		logger.Errorf("Card validation failed: %d erros found", len(err))
		return err
	}

	newClient := &schemas.Client{
		BaseUser: base.BaseUser{
			FirstName: clientRequest.FirstName,
			LastName:  clientRequest.LastName,
			Cpf:       clientRequest.Cpf,
			Email:     clientRequest.Email,
			Phone:     clientRequest.Phone,
		},
		CardData: cardData,
	}

	if err := service.Repository.Create(newClient); err != nil {
		logger.Errorf("Failed to create client in database: %v", err.Error())
		return []*models.ErrorMessage{
			models.CreateErrorMessage("Database", "Failed to create client: "+err.Error()),
		}
	}

	return nil
}

func (service *ClientService) Update(clientRequest *client.UpdateClientRequest) []*models.ErrorMessage {
	logger := config.GetLogger("Update (CLIENT) service")

	if errorMessage := clientRequest.Validate(service.Repository.GetDb()); len(errorMessage) > 0 {
		logger.Errorf("Validation failed: %d errors found", len(errorMessage))
		return errorMessage
	}

	clientDb, err := service.Repository.GetById(clientRequest.Id)
	if err != nil {
		logger.Errorf("Failed to load client from database: %v", err)
		return []*models.ErrorMessage{
			models.CreateErrorMessage("Client", "Error getting client: "+err.Error()),
		}
	}

	if clientRequest.HasCardChanged(clientDb) {
		cardData, cardErrors := clientRequest.ValidateAndFetchCard(config.GetStripeKey())
		if len(cardErrors) > 0 {
			logger.Errorf("Card validation failed: %d errors found", len(cardErrors))
			return cardErrors
		}
		clientDb.CardData = cardData
	}

	clientRequest.MergeInto(clientDb)

	if err := service.Repository.Update(clientDb); err != nil {
		logger.Errorf("Failed to update client in database: %v", err)
		return []*models.ErrorMessage{
			models.CreateErrorMessage("Database", "Failed to update client: "+err.Error()),
		}
	}

	return nil
}

func (service *ClientService) Delete(id uint) *models.ErrorMessage {
	logger := config.GetLogger("Delete (CLIENT) service")

	client, err := service.Repository.GetById(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return &models.ErrorMessage{
			Property: "Client",
			Message:  "Client not found",
		}
	}

	if err != nil {
		errorMessage := fmt.Sprintf("Unexpected error: %v", err.Error())
		logger.Error(errorMessage)
		return &models.ErrorMessage{
			Property: "Database",
			Message:  errorMessage,
		}
	}

	if err := service.Repository.Delete(client); err != nil {
		errorMessage := fmt.Sprintf("Failed to delete client: %v", err.Error())
		logger.Error(errorMessage)
		return &models.ErrorMessage{
			Property: "Database",
			Message:  errorMessage,
		}
	}

	return nil
}

package services

import (
	"github.com/MatheusMikio/dto/client"
	"github.com/MatheusMikio/repository"
)

type IClientService interface {
	GetAll() (*[]client.ClientResponse, error)
	GetById(id uint) (*client.ClientResponse, error)
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

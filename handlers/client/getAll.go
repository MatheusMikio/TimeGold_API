package client

import (
	"net/http"

	"github.com/MatheusMikio/dto/client"
	"github.com/MatheusMikio/handlers"
	"github.com/MatheusMikio/schemas"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func GetAllHandler(service services.IClientService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clients := []schemas.Client{}

		if err := handlers.Db.Preload("Appointments").Find(&clients).Error; err != nil {
			handlers.SendError(ctx, http.StatusInternalServerError, "error listing clients")
			return
		}

		clientsResponse := []client.ClientResponse{}

		for _, cli := range clients {
			appointmentsSummary := []client.AppointmentSummary{}

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
		handlers.SendSuccess(ctx, http.StatusOK, "GET Clients", &clientsResponse)
	}
}

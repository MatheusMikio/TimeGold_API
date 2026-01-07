package client

import (
	"net/http"

	"github.com/MatheusMikio/handlers"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func GetAllHandler(service services.IClientService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clients, err := service.GetAll()

		if err != nil {
			handlers.SendError(ctx, http.StatusInternalServerError, err.Error())
			return

		}
		handlers.SendSuccess(ctx, http.StatusOK, "GET Clients", &clients)
	}
}

package client

import (
	"fmt"
	"net/http"

	"github.com/MatheusMikio/config"
	"github.com/MatheusMikio/handlers"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func GetAllHandler(service services.IClientService) gin.HandlerFunc {
	logger := config.GetLogger("GetAll (CLIENT) handler")
	return func(ctx *gin.Context) {
		clients, err := service.GetAll()

		if err != nil {
			errorMessage := fmt.Sprintf("Error listing clients: %v", err.Error())
			logger.Error(errorMessage)
			handlers.SendError(ctx, http.StatusInternalServerError, errorMessage)
			return
		}

		handlers.SendSuccess(ctx, http.StatusOK, "GET Clients", *clients)
	}
}

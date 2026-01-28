package client

import (
	"net/http"

	"github.com/MatheusMikio/config"
	"github.com/MatheusMikio/dto/client"
	"github.com/MatheusMikio/handlers"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func UpdateHandler(service services.IClientService) gin.HandlerFunc {
	logger := config.GetLogger("Update (CLIENT) handler")
	return func(ctx *gin.Context) {
		request := &client.UpdateClientRequest{}

		if err := ctx.BindJSON(&request); err != nil {
			logger.Errorf("Error binding JSON: %v", err)
			handlers.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		if errorMessage := service.Update(request); len(errorMessage) > 0 {
			handlers.SendError(ctx, http.StatusBadRequest, errorMessage)
			return
		}

		handlers.SendSuccess(ctx, http.StatusOK, "PUT Client", "Client Updated")
	}
}

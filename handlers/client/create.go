package client

import (
	"net/http"

	"github.com/MatheusMikio/config"
	"github.com/MatheusMikio/dto/client"
	"github.com/MatheusMikio/handlers"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func CreateHandler(service services.IClientService) gin.HandlerFunc {
	logger := config.GetLogger("Create (CLIENT)")
	return func(ctx *gin.Context) {
		request := &client.ClientRequest{}

		if err := ctx.BindJSON(&request); err != nil {
			logger.Errorf("Error binding JSON: %v", err)
			handlers.SendError(ctx, http.StatusBadRequest, err)
			return
		}

		errorMessage := service.Create(request)
		if len(errorMessage) > 0 {
			handlers.SendError(ctx, http.StatusBadRequest, errorMessage)
			return
		}

		handlers.SendSuccess(ctx, http.StatusCreated, "POST Client", "Client Created")
	}
}

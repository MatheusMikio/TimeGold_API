package client

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MatheusMikio/config"
	"github.com/MatheusMikio/handlers"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func GetByIdHandler(service services.IClientService) gin.HandlerFunc {
	logger := config.GetLogger("Get (CLIENT) handler")
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)

		if id <= 0 {
			handlers.SendError(ctx, http.StatusBadRequest, "Invalid client ID: The ID must be greater than 0.")
			return
		}

		if err != nil {
			errorMessage := fmt.Sprintf("Invalid client ID: %v", err.Error())
			logger.Error(errorMessage)
			handlers.SendError(ctx, http.StatusBadRequest, errorMessage)
			return
		}

		client, err := service.GetById(uint(id))

		if err != nil {
			errorMessage := fmt.Sprintf("Error get client: %v", err.Error())
			logger.Error(errorMessage)
			handlers.SendError(ctx, http.StatusNotFound, errorMessage)
			return
		}

		handlers.SendSuccess(ctx, http.StatusOK, "GET Client", *client)

	}
}

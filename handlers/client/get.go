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

func GetHandler(service services.IClientService) gin.HandlerFunc {
	logger := config.GetLogger("Get (Client)")
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			errorMessage := "Invalid client ID"
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

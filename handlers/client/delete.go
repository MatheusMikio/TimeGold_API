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

func DeleteHandler(service services.IClientService) gin.HandlerFunc {
	logger := config.GetLogger("Delete (Client) handler")
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

		if err := service.Delete(uint(id)); err != nil {
			if err.Property == "Client" && err.Message == "Client not found" {
				handlers.SendError(ctx, http.StatusNotFound, err)
				return
			}
			logger.Error(err.Message)
			handlers.SendError(ctx, http.StatusInternalServerError, err)
			return
		}

		handlers.SendSuccess(ctx, http.StatusNoContent, "DELETE Client", "Client Deleted")
	}
}

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

const (
	maxSize uint64 = 100
)

func GetAllHandler(service services.IClientService) gin.HandlerFunc {
	logger := config.GetLogger("GetAll (CLIENT) handler")
	return func(ctx *gin.Context) {
		pageStr := ctx.DefaultQuery("page", "1")
		sizeStr := ctx.DefaultQuery("size", "12")
		page, err := strconv.ParseUint(pageStr, 10, 64)

		if err != nil || page == 0 {
			handlers.SendError(ctx, http.StatusBadRequest, `Invalid page parameter (must be greater than 0)`)
			return
		}

		size, err := strconv.ParseUint(sizeStr, 10, 64)
		if err != nil || size == 0 || size > maxSize {
			handlers.SendError(ctx, http.StatusBadRequest, `Invalid "size" parameter (must be greater than 0 and at most 100)`)
			return
		}

		clients, err := service.GetAll(page, size)

		if err != nil {
			errorMessage := fmt.Sprintf("Error listing clients: %v", err.Error())
			logger.Error(errorMessage)
			handlers.SendError(ctx, http.StatusInternalServerError, errorMessage)
			return
		}

		handlers.SendSuccess(ctx, http.StatusOK, "GET Clients", *clients)
	}
}

package client

import (
	"net/http"

	"github.com/MatheusMikio/handlers"
	"github.com/MatheusMikio/schemas"
	"github.com/gin-gonic/gin"
)

func GetAllHandler(ctx *gin.Context) {
	clients := []schemas.Client{}
	if err := handlers.Db.Find(&clients).Error; err != nil {
		handlers.SendError(ctx, http.StatusInternalServerError, "error listing clients")
		return
	}
	handlers.SendSuccess(ctx, http.StatusOK, "GET Clients", &clients)
}

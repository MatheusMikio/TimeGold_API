package auth

import (
	"net/http"

	"github.com/MatheusMikio/handlers"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func ClientGoogleAuth(service services.IAuthClientService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state := generateState()
		url := service.GetGoogleAuthUrl(state)

		handlers.SendSuccess(ctx, http.StatusOK, "ClientGoogleAuth", map[string]string{"url": url})
	}
}

func ClientGoogleCallBack(service services.IAuthClientService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Query("code")

		if code == "" {
			handlers.SendError(ctx, http.StatusBadRequest, "Authorization code is required")
			return
		}

		response, err := service.HandleGoogleCallBack(code)
		if err != nil {
			handlers.SendError(ctx, http.StatusUnauthorized, err)
			return
		}

		handlers.SendAuthSuccess(ctx, http.StatusOK, response.Token, 7200, response.User)
	}
}

package auth

import (
	"net/http"

	"github.com/MatheusMikio/handlers"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func ProfessionalGoogleAuth(service services.IAuthProfessionalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state := generateState()
		url := service.GetGoogleAuthUrl(state)

		handlers.SendSuccess(ctx, http.StatusOK, "ProfessionalGoogleAuth", map[string]string{"url": url})
	}
}

func ProfessionalGoogleCallback(service services.IAuthProfessionalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Query("code")

		if code == "" {
			handlers.SendError(ctx, http.StatusBadRequest, "Authorization code is required")
			return
		}

		response, errMsg := service.HandleGoogleCallBack(code)
		if errMsg != nil {
			handlers.SendError(ctx, http.StatusUnauthorized, errMsg)
			return
		}

		handlers.SendAuthSuccess(ctx, http.StatusOK, response.Token, 7200, response.User)
	}
}

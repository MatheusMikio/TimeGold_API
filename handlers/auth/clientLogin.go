package auth

import (
	"net/http"

	"github.com/MatheusMikio/dto/auth"
	"github.com/MatheusMikio/handlers"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func ClientRequestMagicLink(service services.IAuthClientService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := auth.LoginRequest{}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			handlers.SendError(ctx, http.StatusBadRequest, "Invalid email format")
			return
		}

		if errMsg := service.RequestMagicLink(request.Email); errMsg != nil {
			handlers.SendError(ctx, http.StatusBadRequest, errMsg)
			return
		}
		handlers.SendSuccess(ctx, http.StatusOK, "ClientRequestMagicLink", "Magic link sent to your email")
	}
}

func ClientVerifyMagicLink(service services.IAuthClientService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query("token")
		if token == "" {
			handlers.SendError(ctx, http.StatusBadRequest, "Token is required")
			return
		}

		response, err := service.VerifyMagicLink(token)
		if err != nil {
			handlers.SendError(ctx, http.StatusBadRequest, err)
			return
		}

		handlers.SendAuthSuccess(ctx, http.StatusOK, response.Token, 7200, response.User)
	}

}

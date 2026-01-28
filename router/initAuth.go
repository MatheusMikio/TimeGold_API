package router

import (
	ah "github.com/MatheusMikio/handlers/auth"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func initAuth(rg *gin.RouterGroup, authClientService services.IAuthClientService, authProfessionalService services.IAuthProfessionalService) {
	auth := rg.Group("/auth")
	{
		client := auth.Group("/client")
		{
			client.POST("/login", ah.ClientRequestMagicLink(authClientService))
			client.GET("/verify", ah.ClientVerifyMagicLink(authClientService))
			client.GET("/google", ah.ClientGoogleAuth(authClientService))
			client.GET("/google/callback", ah.ClientGoogleCallBack(authClientService))
		}

		professional := auth.Group("/professional")
		{
			professional.POST("/login", ah.ProfessionalRequestMagicLink(authProfessionalService))
			professional.GET("/verify", ah.ProfessionalVerifyMagicLink(authProfessionalService))
			professional.GET("/google", ah.ProfessionalGoogleAuth(authProfessionalService))
			professional.GET("/google/callback", ah.ProfessionalGoogleCallback(authProfessionalService))
		}
	}
}

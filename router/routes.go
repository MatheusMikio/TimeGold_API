package router

import (
	"github.com/MatheusMikio/middlewares"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func initRoutes(
	router *gin.Engine,
	clientService services.IClientService,
	//Auth
	authClientService services.IAuthClientService,
	authProfessionalService services.IAuthProfessionalService,
) {
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		initAuth(v1, authClientService, authProfessionalService)

		initCompany(v1)
		initClient(v1, clientService)
		initProfessional(v1)

		protected := v1.Group("")
		protected.Use(middlewares.AuthRequired())
		{
			initService(protected)
			initScheduling(protected)
		}

	}

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(files.handler))
}

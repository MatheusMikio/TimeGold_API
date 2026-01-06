package router

import (
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func initRoutes(
	router *gin.Engine,
	clientService services.IClientService,
) {
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		initCompany(v1)
		initClient(v1, clientService)
		initProfessional(v1)
		initService(v1)
		initScheduling(v1)
	}

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(files.handler))
}

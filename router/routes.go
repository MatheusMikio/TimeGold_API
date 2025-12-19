package router

import (
	"github.com/MatheusMikio/handlers"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	handlers.Init()
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		initCompany(v1)
		initClient(v1)
		initProfessional(v1)
		initService(v1)
		initScheduling(v1)
	}
}

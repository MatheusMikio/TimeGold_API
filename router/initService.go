package router

import (
	sh "github.com/MatheusMikio/handlers/service"
	"github.com/gin-gonic/gin"
)

func initService(rg *gin.RouterGroup) {
	service := rg.Group("/services")
	{
		service.GET("", sh.GetAllHandler)
		service.GET("/:id", sh.GetByIdHandler)
		service.POST("", sh.CreateHandler)
		service.PUT("", sh.UpdateHandler)
		service.DELETE("/:id", sh.DeleteHandler)
	}
}

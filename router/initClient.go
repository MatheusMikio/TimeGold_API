package router

import (
	ch "github.com/MatheusMikio/handlers/client"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func initClient(rg *gin.RouterGroup, service services.IClientService) {
	client := rg.Group("/clients")
	{
		client.GET("", ch.GetAllHandler(service))
		client.GET("/:id", ch.GetHandler)
		client.POST("", ch.CreateHandler)
		client.PUT("/:id", ch.UpdateHandler)
		client.DELETE("/:id", ch.DeleteHandler)
	}
}

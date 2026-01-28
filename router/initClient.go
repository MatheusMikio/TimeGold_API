package router

import (
	ch "github.com/MatheusMikio/handlers/client"
	"github.com/MatheusMikio/middlewares"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
)

func initClient(rg *gin.RouterGroup, service services.IClientService) {
	client := rg.Group("/client")
	{
		client.POST("", ch.CreateHandler(service))
	}

	clientProtected := rg.Group("/client")
	clientProtected.Use(middlewares.AuthRequired())
	// clientProtected.Use(middlewares.RoleRequired("ADMIN"))
	{
		clientProtected.GET("", ch.GetAllHandler(service))
		clientProtected.GET("/:id", ch.GetByIdHandler(service))
		clientProtected.PUT("", ch.UpdateHandler(service))
		clientProtected.DELETE("/:id", ch.DeleteHandler(service))
	}
}

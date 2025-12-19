package router

import (
	ch "github.com/MatheusMikio/handlers/client"
	"github.com/gin-gonic/gin"
)

func initClient(rg *gin.RouterGroup) {
	client := rg.Group("/clients")
	{
		client.GET("", ch.GetAllHandler)
		client.GET("", ch.GetHandler)
		client.POST("", ch.CreateHandler)
		client.PUT("", ch.UpdateHandler)
		client.DELETE("", ch.DeleteHandler)
	}
}

package router

import (
	ph "github.com/MatheusMikio/handlers/professional"
	"github.com/gin-gonic/gin"
)

func initProfessional(rg *gin.RouterGroup) {
	professional := rg.Group("/professionals")
	{
		professional.GET("", ph.GetAllHandler)
		professional.GET("/:id", ph.GetHandler)
		professional.POST("", ph.CreateHandler)
		professional.PUT("/:id", ph.UpdateHandler)
		professional.DELETE("/:id", ph.DeleteHandler)
	}
}

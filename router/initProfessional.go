package router

import (
	ph "github.com/MatheusMikio/handlers/professional"
	"github.com/MatheusMikio/middlewares"
	"github.com/gin-gonic/gin"
)

func initProfessional(rg *gin.RouterGroup) {
	professional := rg.Group("/professionals")
	professional.Use(middlewares.RoleRequired("PROFESSIONAL", "ADMIN"))
	{
		professional.GET("", ph.GetAllHandler)
		professional.GET("/:id", ph.GetByIdHandler)
		professional.POST("", ph.CreateHandler)
		professional.PUT("", ph.UpdateHandler)
		professional.DELETE("/:id", ph.DeleteHandler)
	}
}

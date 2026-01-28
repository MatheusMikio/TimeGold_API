package router

import (
	ch "github.com/MatheusMikio/handlers/company"
	"github.com/MatheusMikio/middlewares"
	"github.com/gin-gonic/gin"
)

func initCompany(rg *gin.RouterGroup) {
	company := rg.Group("/companies")
	company.Use(middlewares.RoleRequired("ADMIN"))
	{
		company.GET("", ch.GetAllHandler)
		company.GET("/:id", ch.GetByIdHandler)
		company.POST("", ch.CreateHandler)
		company.PUT("", ch.UpdateHandler)
		company.DELETE("/:id", ch.DeleteHandler)
	}
}

package router

import (
	ch "github.com/MatheusMikio/handlers/company"
	"github.com/gin-gonic/gin"
)

func initCompany(rg *gin.RouterGroup) {
	company := rg.Group("/companies")
	{
		company.GET("", ch.GetAllHandler)
		company.GET("", ch.GetHandler)
		company.POST("", ch.CreateHandler)
		company.PUT("", ch.UpdateHandler)
		company.DELETE("", ch.DeleteHandler)
	}
}

package router

import "github.com/gin-gonic/gin"

func initCompany(rg *gin.RouterGroup) {
	company := rg.Group("/companies")
	{
		company.POST("", handlers.CreateCompany)

	}
}

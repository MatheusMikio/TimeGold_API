package router

import "github.com/gin-gonic/gin"

func initCompany(rg *gin.RouterGroup) {
	company := rg.Group("/companies")
	{
		company.GET("", company.GetAllHandler)
		company.GET("", company.GetHandler)
		company.POST("", company.CreateHandler)
		company.PUT("", company.UpdateHandler)
		company.DELETE("", company.DeleteHandler)
	}
}

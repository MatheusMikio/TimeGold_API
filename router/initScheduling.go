package router

import (
	sh "github.com/MatheusMikio/handlers/scheduling"
	"github.com/gin-gonic/gin"
)

func initScheduling(rg *gin.RouterGroup) {
	scheduling := rg.Group("/schedulings")
	{
		scheduling.GET("", sh.GetAllHandler)
		scheduling.GET("", sh.GetHandler)
		scheduling.POST("", sh.CreateHandler)
		scheduling.PUT("", sh.UpdateHandler)
		scheduling.DELETE("", sh.DeleteHandler)
	}
}

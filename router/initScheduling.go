package router

import (
	sh "github.com/MatheusMikio/handlers/scheduling"
	"github.com/gin-gonic/gin"
)

func initScheduling(rg *gin.RouterGroup) {
	scheduling := rg.Group("/schedulings")
	{
		scheduling.GET("", sh.GetAllHandler)
		scheduling.GET("/:id", sh.GetHandler)
		scheduling.POST("", sh.CreateHandler)
		scheduling.PUT("/:id", sh.UpdateHandler)
		scheduling.DELETE("/:id", sh.DeleteHandler)
	}
}

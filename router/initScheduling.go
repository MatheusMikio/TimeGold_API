package router

func initScheduling(rg *gin.RouterGroup){
	scheduling := rg.Group("/schedulings")
	{
		scheduling.GET("", scheduling.GetAllHandler)
		scheduling.GET("", scheduling.GetHandler)
		scheduling.POST("", scheduling.CreateHandler)
		scheduling.PUT("", scheduling.UpdateHandler)
		scheduling.DELETE("", scheduling.DeleteHandler)
	}
}
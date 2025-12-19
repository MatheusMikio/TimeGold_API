package router

func initService(rg *gin.RouterGroup){
	service := rg.Group("/services")
	{
		service.GET("", service.GetAllHandler)
		service.GET("", service.GetHandler)
		service.POST("", service.CreateHandler)
		service.PUT("", service.UpdateHandler)
		service.DELETE("", service.DeleteHandler)
	}
}
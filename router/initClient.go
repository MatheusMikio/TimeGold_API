package router

func initClient(rg *gin.RouterGroup){
	client := rg.Group("/clients")
	{
		client.GET("", client.GetAllHandler)
		client.GET("", client.GetHandler)
		client.POST("", client.CreateHandler)
		client.PUT("", client.UpdateHandler)
		client.DELETE("", client.DeleteHandler)
	}
}
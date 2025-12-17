package router

func initClient(rg *gin.RouterGroup){
	client := rg.Group("/clients")
	{
		client.POST("", CreateClientHandler)
	}
}
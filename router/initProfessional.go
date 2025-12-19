package router

func initProfessional(rg *gin.RouterGroup){
	professional := rg.Group("/professionals")
	{
		professional.GET("", professional.GetAllHandler)
		professional.GET("", professional.GetHandler)
		professional.POST("", professional.CreateHandler)
		professional.PUT("", professional.UpdateHandler)
		professional.DELETE("", professional.DeleteHandler)
	}
}
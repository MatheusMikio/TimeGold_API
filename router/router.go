package router

import (
	"github.com/MatheusMikio/handlers"
	"github.com/MatheusMikio/repository"
	"github.com/MatheusMikio/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	handlers.Init()

	clientRepo := repository.NewClientRepository(db)

	clientService := services.NewClientService(clientRepo)

	r := gin.Default()

	initRoutes(r,
		clientService,
	)

	r.Run(":8080")
}

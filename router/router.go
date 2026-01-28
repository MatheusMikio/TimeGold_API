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
	professionalRepo := repository.NewProfessionalRepository(db)
	magicLinkRepo := repository.NewMagicLinkRepository(db)

	clientService := services.NewClientService(clientRepo)

	// Auth
	authClientService := services.NewAuthClientService(clientRepo, magicLinkRepo)
	authProfessionalService := services.NewAuthProfessionalService(professionalRepo, magicLinkRepo)

	r := gin.Default()

	initRoutes(r,
		clientService,
		authClientService,
		authProfessionalService,
	)

	r.Run(":8080")
}

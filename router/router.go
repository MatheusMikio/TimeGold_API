package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	initRoutes(r)
	r.Run(":8080")
}

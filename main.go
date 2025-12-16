package main

import (
	"github.com/MatheusMikio/config"
	"github.com/MatheusMikio/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
	}
	router.Init()
}

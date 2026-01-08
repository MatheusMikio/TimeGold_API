package handlers

import (
	"github.com/MatheusMikio/config"
)

var (
	Logger *config.Logger
)

func Init() {
	Logger = config.GetLogger("handler")
}

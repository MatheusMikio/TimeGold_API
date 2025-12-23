package handlers

import (
	"github.com/MatheusMikio/config"
	"gorm.io/gorm"
)

var (
	Logger *config.Logger
	Db     *gorm.DB
)

func Init() {
	Logger = config.GetLogger("handler")
	Db = config.GetDb()
}

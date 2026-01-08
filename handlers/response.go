package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, msg interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, code int, op string, data ...interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("Operation from handler %s successfull", op),
		"data":    data,
	})
}

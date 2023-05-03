package routes

import (
	"awesomeProject/app/serverCore"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", serverCore.PingHander)

	return r
}

package routes

import (
	"awesomeProject/app/serverCore/handles"
	"awesomeProject/app/serverCore/routes"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", handles.PingHander)
	routes.LoadServerCoreRouter(r)
	return r
}

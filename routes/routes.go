package routes

import (
	"awesomeProject/app/cache"
	"awesomeProject/app/serverCore"
	"awesomeProject/app/serverCore/handles"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", handles.PingHander)
	serverCore.LoadServerCoreRouter(r)
	cache.LoadCacheRouter(r)
	return r
}

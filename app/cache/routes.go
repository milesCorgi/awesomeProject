package cache

import (
	"awesomeProject/app/cache/handles"
	"github.com/gin-gonic/gin"
)

func LoadCacheRouter(r *gin.Engine) {
	r.POST("/api/cache/getAllCache", handles.GetAllCache)
	r.POST("/api/cache/getACache", handles.GetACache)
	r.POST("/api/cache/setACache", handles.SetACache)
	r.POST("/api/cache/clearCache", handles.ClearCache)
}

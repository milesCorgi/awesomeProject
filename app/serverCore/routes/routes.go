package routes

import (
	"awesomeProject/app/serverCore"
	"github.com/gin-gonic/gin"
)

func LoadServerCoreRouter(r *gin.Engine) {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r.POST("/api/query_two_set_video_info", serverCore.ShowVideoInfo)
	r.POST("/api/query_youtube_video_list", serverCore.ShowYoutubeVideoList)
	r.POST("/api/query_bilibili_video_list", serverCore.ShowVideoInfo)

}

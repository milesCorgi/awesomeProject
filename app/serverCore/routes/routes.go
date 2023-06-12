package routes

import (
	"awesomeProject/app/serverCore/handles"
	"github.com/gin-gonic/gin"
)

func LoadServerCoreRouter(r *gin.Engine) {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r.POST("/api/query_two_set_video_info", handles.ShowVideoInfo)
	r.POST("/api/query_youtube_video_list", handles.ShowYoutubeVideoList)
	r.POST("/api/query_bilibili_video_list", handles.ShowBiliBiliVideoList)

	r.POST("/api/get_keywords", handles.GetKeyWord)
}

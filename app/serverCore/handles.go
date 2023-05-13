package serverCore

import (
	"awesomeProject/app/serverCore/dto"
	"awesomeProject/middleware/db"
	"awesomeProject/middleware/log"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHander(c *gin.Context) {
	c.String(200, "pong")
}

func DBQueryDemo(c *gin.Context) {
	var videoNoteInfos []models.VideoNoteInfo
	var queryAInfo dto.QueryAInfo
	bindJSONErr := c.ShouldBindJSON(&queryAInfo)
	if bindJSONErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindJSONErr.Error()})
		return
	}
	log.Logger.Info("Info------------" + queryAInfo.Info)
	log.Logger.Info("BiliBiliID------------" + queryAInfo.BiliBiliID)
	log.Logger.Info("YoutubeLink------------" + queryAInfo.YoutubeLink)
	log.Logger.Info("OtherLink------------" + queryAInfo.OtherLink)
	log.Logger.Info("InfoType------------" + queryAInfo.InfoType)
	log.Logger.Info("KeyWord------------" + queryAInfo.KeyWord)
	queryTable := map[string]interface{}{}

	if queryAInfo.BiliBiliID != "" {
		queryTable["bili_bili_id"] = queryAInfo.BiliBiliID
	}
	if queryAInfo.BiliBiliID != "" {
		queryTable["youtube_link"] = queryAInfo.YoutubeLink
	}
	if queryAInfo.OtherLink != "" {
		queryTable["other_link"] = queryAInfo.OtherLink
	}
	if queryAInfo.InfoType != "" {
		queryTable["info_type"] = queryAInfo.InfoType
	}
	if queryAInfo.KeyWord != "" {
		queryTable["keyword"] = queryAInfo.KeyWord
	}
	db.DB = db.DB.Limit(200).Where(queryTable)
	// 读取所有数据
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "获取萌点失败，请重试或者联系狗狗"}
	if queryAInfo.Info != "" {
		// todo: 想办法搞成子查询
		db.DB = db.DB.Limit(200).Where("info = ?", queryAInfo.Info)
	}
	err := db.DB.Limit(200).Find(&videoNoteInfos).Error
	// Limit得放前面才生效

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wholeResult["total_num"] = len(videoNoteInfos)
	wholeResult["info_list"] = videoNoteInfos
	wholeResult["msg"] = "success"
	wholeResult["error_num"] = 0
	c.JSON(http.StatusOK, wholeResult)
}

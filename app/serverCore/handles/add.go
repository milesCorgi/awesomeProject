package handles

import (
	"awesomeProject/app/serverCore/dto"
	"awesomeProject/middleware/db"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var cstSh, _ = time.LoadLocation("Asia/Shanghai")

func AddVideoInfo(c *gin.Context) {
	var videoNoteInfos []models.VideoNoteInfo
	var QueryVideoInfo dto.VideoInfo
	bindJSONErr := c.ShouldBindJSON(&QueryVideoInfo)
	db := db.DB
	// 过滤器得初始化才行
	db = db.Limit(10000)
	if bindJSONErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindJSONErr.Error()})
		return
	}
	// 目前暂定描述和所属ip完全一致的萌点属于同一条萌点
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "获取萌点失败，请重试或者联系狗狗"}
	queryTable := map[string]interface{}{}
	queryTable["info"] = QueryVideoInfo.Info
	queryTable["intellectual_property_name"] = QueryVideoInfo.IntellectualPropertyName
	results := db.Where(queryTable).Find(&videoNoteInfos).Error
	if results != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error()})
		return
	}
	if len(videoNoteInfos) == 0 {
		var videoNoteInfo = models.VideoNoteInfo{
			KeyWord:                  QueryVideoInfo.KeyWord,
			BiliBiliLink:             QueryVideoInfo.BiliBiliLink,
			BiliBiliID:               QueryVideoInfo.BiliBiliID,
			YoutubeLink:              QueryVideoInfo.YoutubeLink,
			OtherLink:                QueryVideoInfo.OtherLink,
			ImgLinks:                 QueryVideoInfo.ImgLinks,
			FanFictionLink:           QueryVideoInfo.FanFictionLink,
			InfoType:                 QueryVideoInfo.InfoType,
			EditPassWord:             QueryVideoInfo.EditPassWord,
			Enable:                   1,
			User:                     QueryVideoInfo.User,
			AddTime:                  time.Now().In(cstSh).Format("2006-01-02 15:04:05"),
			UpdateTime:               time.Now().In(cstSh).Format("2006-01-02 15:04:05"),
			IntellectualPropertyName: QueryVideoInfo.IntellectualPropertyName,
			Info:                     QueryVideoInfo.Info,
		}
		resultsAdd := db.Debug().Create(&videoNoteInfo)
		if resultsAdd.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultsAdd})
			return
		}
	} else {
		resultsUpdate := db.Model(&videoNoteInfos).
			Debug().
			Where(queryTable).
			Updates(map[string]interface{}{
				"key_word":         QueryVideoInfo.KeyWord,
				"bili_bili_link":   QueryVideoInfo.BiliBiliLink,
				"bili_bili_id":     QueryVideoInfo.BiliBiliID,
				"youtube_link":     QueryVideoInfo.YoutubeLink,
				"other_link":       QueryVideoInfo.OtherLink,
				"img_links":        QueryVideoInfo.ImgLinks,
				"fan_fiction_link": QueryVideoInfo.FanFictionLink,
				"info_type":        QueryVideoInfo.InfoType,
				"edit_pass_word":   QueryVideoInfo.EditPassWord,
				"enable":           1,
				"user":             QueryVideoInfo.User,
				"add_time":         time.Now().In(cstSh).Format("2006-01-02 15:04:05"),
				"update_time":      time.Now().In(cstSh).Format("2006-01-02 15:04:05"),
			})
		if resultsUpdate.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultsUpdate.Error})
			return
		}
	}
	wholeResult["total_num"] = len(videoNoteInfos)
	wholeResult["info_list"] = videoNoteInfos
	wholeResult["msg"] = "success"
	wholeResult["error_num"] = 0
	c.JSON(http.StatusOK, wholeResult)
}

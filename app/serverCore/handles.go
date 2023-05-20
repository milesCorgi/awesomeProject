package serverCore

import (
	"awesomeProject/app/serverCore/dto"
	"awesomeProject/middleware/db"
	"awesomeProject/middleware/log"
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHander(c *gin.Context) {
	c.String(200, "pong")
}

func ShowVideoInfo(c *gin.Context) {
	var videoNoteInfos []models.VideoNoteInfo
	var QueryVideoInfo dto.QueryVideoInfo
	bindJSONErr := c.ShouldBindJSON(&QueryVideoInfo)
	db := db.DB
	// 过滤器得初始化才行
	db = db.Limit(10000)
	if bindJSONErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindJSONErr.Error()})
		return
	}

	log.Logger.Info(fmt.Sprintf("ShowVideoInfo req------------%v", QueryVideoInfo))
	queryTable := map[string]interface{}{}

	if QueryVideoInfo.BiliBiliID != "" {
		queryTable["bili_bili_id"] = QueryVideoInfo.BiliBiliID
	}
	if QueryVideoInfo.BiliBiliID != "" {
		queryTable["youtube_link"] = QueryVideoInfo.YoutubeLink
	}
	if QueryVideoInfo.OtherLink != "" {
		queryTable["other_link"] = QueryVideoInfo.OtherLink
	}
	if QueryVideoInfo.InfoType != "" {
		queryTable["info_type"] = QueryVideoInfo.InfoType
	}
	if QueryVideoInfo.KeyWord != "" {
		queryTable["keyword"] = QueryVideoInfo.KeyWord
	}
	db = db.Where(queryTable)
	// 读取所有数据
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "获取萌点失败，请重试或者联系狗狗"}
	if QueryVideoInfo.Info != "" {
		// 模糊匹配目前就找到了这个法子（好糙）
		db = db.
			Where("info like ? ", QueryVideoInfo.Info+"%").
			Or("info like ? ", "%"+QueryVideoInfo.Info).
			Or("info like ? ", "%"+QueryVideoInfo.Info+"%")
	}
	err := db.Find(&videoNoteInfos).Error
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

func ShowYoutubeVideoList(c *gin.Context) {

	var QueryWebVideoInfo dto.QueryWebVideoInfo
	bindJSONErr := c.ShouldBindJSON(&QueryWebVideoInfo)
	if bindJSONErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindJSONErr.Error()})
		return
	}
	QueryWebVideoInfo.WebName = "youtube"
	log.Logger.Info(fmt.Sprintf("ShowYoutubeVideoList req------------%v", QueryWebVideoInfo))
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "获取萌点失败，请重试或者联系狗狗"}
	wholeResult = ShowWebVideoList(QueryWebVideoInfo)
	c.JSON(http.StatusOK, wholeResult)
}

func ShowWebVideoList(QueryWebVideoInfo dto.QueryWebVideoInfo) map[string]interface{} {

	// 这里需要注意一个细节,首先将全局的db变量赋值给了Db,如果用db直接进行操作,那一系列的赋值语句将会影响db的地址,影响后续的数据库操作.
	// https://cloud.tencent.com/developer/article/1374623
	db := db.DB
	db.Limit(10000)
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "获取萌点失败，请重试或者联系狗狗"}
	// todo:做成分表
	var YoutubeVideoList []models.YoutubeVideoList
	log.Logger.Info(fmt.Sprintf("ShowWebVideoList req------------%v", QueryWebVideoInfo))
	if QueryWebVideoInfo.From == "" && QueryWebVideoInfo.To == "" && QueryWebVideoInfo.Keyword == "" {
		OrderBy := "published_at"
		if QueryWebVideoInfo.IfDesc == "1" {
			OrderBy = "published_at desc"
		} else {
			OrderBy = "published_at"
		}
		db = db.Order(OrderBy)

		err := db.Debug().Find(&YoutubeVideoList).Error
		// Limit得放前面才生效
		if err != nil {
			log.Logger.Error(err.Error())
			return wholeResult
		} else {
			//for _, Video := range TwoSetVideoInfos {
			//	logs.Debug(Video)
			//}
			wholeResult["total_num"] = len(YoutubeVideoList)
			wholeResult["info_list"] = YoutubeVideoList
			wholeResult["msg"] = "success"
			wholeResult["error_num"] = 0
		}
	} else {
		log.Logger.Debug("开始查询")
		OrderBy := "published_at"
		if QueryWebVideoInfo.From != "" {
			db = db.Where("published_at > ?", QueryWebVideoInfo.From).
				Or("published_at = ?", QueryWebVideoInfo.From)
		}
		if QueryWebVideoInfo.To != "" {
			db = db.Where("published_at < ?", QueryWebVideoInfo.To).
				Or("published_at = ?", QueryWebVideoInfo.To)
		}
		if QueryWebVideoInfo.Keyword != "" {
			db = db.
				Where("title like ? ", QueryWebVideoInfo.Keyword+"%").
				Or("title like ? ", "%"+QueryWebVideoInfo.Keyword).
				Or("title like ? ", "%"+QueryWebVideoInfo.Keyword+"%")
		}
		if QueryWebVideoInfo.IfDesc == "1" {
			OrderBy = "-published_at"
		} else {
			OrderBy = "published_at"
		}

		db = db.Order(OrderBy)
		err := db.Debug().Find(&YoutubeVideoList).Error
		if err != nil {
			log.Logger.Error(err.Error())
			return wholeResult
		} else {
			//for _, info := range TwoSetVideoInfos {
			//	logs.Info(info)
			//}
			wholeResult["total_num"] = len(YoutubeVideoList)
			wholeResult["info_list"] = YoutubeVideoList
			wholeResult["msg"] = "success"
			wholeResult["error_num"] = 0
		}
	}
	log.Logger.Info(fmt.Sprintf("ShowYoutubeVideoList resp------------%v", wholeResult))
	return wholeResult
}

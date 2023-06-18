package handles

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
	var QueryVideoInfo dto.VideoInfo
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
	queryTable["enable"] = "1"
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
		queryTable["key_word"] = QueryVideoInfo.KeyWord
	}
	if QueryVideoInfo.IntellectualPropertyName != "" {
		queryTable["intellectual_property_name"] = QueryVideoInfo.IntellectualPropertyName
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

var SourceMap = map[string]int{"bilibili": 1, "youtube": 2}

func ShowBiliBiliVideoList(c *gin.Context) {
	var QueryWebVideoInfo dto.QueryWebVideoInfo
	bindJSONErr := c.ShouldBindJSON(&QueryWebVideoInfo)
	if bindJSONErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindJSONErr.Error()})
		return
	}
	QueryWebVideoInfo.SourceId = SourceMap["bilibili"]
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "获取萌点失败，请重试或者联系狗狗"}
	wholeResult = ShowVideoList(QueryWebVideoInfo)
	c.JSON(http.StatusOK, wholeResult)
}

func ShowYoutubeVideoList(c *gin.Context) {
	var QueryWebVideoInfo dto.QueryWebVideoInfo
	bindJSONErr := c.ShouldBindJSON(&QueryWebVideoInfo)
	if bindJSONErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindJSONErr.Error()})
		return
	}
	QueryWebVideoInfo.SourceId = SourceMap["youtube"]
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "获取萌点失败，请重试或者联系狗狗"}
	wholeResult = ShowVideoList(QueryWebVideoInfo)
	c.JSON(http.StatusOK, wholeResult)
}

func ShowVideoList(QueryWebVideoInfo dto.QueryWebVideoInfo) map[string]interface{} {

	// 这里需要注意一个细节,首先将全局的db变量赋值给了Db,如果用db直接进行操作,那一系列的赋值语句将会影响db的地址,影响后续的数据库操作.
	// https://cloud.tencent.com/developer/article/1374623
	dbShowVideoList := db.DB
	dbShowVideoList.Limit(10000)
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "获取萌点失败，请重试或者联系狗狗"}
	// todo:做成分表
	var VideoList []models.VideoList
	dbShowVideoList = dbShowVideoList.Debug().Where("source_id", int64(QueryWebVideoInfo.SourceId))
	log.Logger.Debug("开始查询")
	OrderBy := "published_at"
	if QueryWebVideoInfo.IfDesc == "1" {
		OrderBy = "published_at desc"
	}
	dbShowVideoList = dbShowVideoList.Order(OrderBy)
	if QueryWebVideoInfo.Keyword != "" {
		dbShowVideoList = dbShowVideoList.
			Where("title like ? ", QueryWebVideoInfo.Keyword+"%").
			Or("title like ? ", "%"+QueryWebVideoInfo.Keyword).
			Or("title like ? ", "%"+QueryWebVideoInfo.Keyword+"%")
	}

	if QueryWebVideoInfo.From != "" {
		dbShowVideoList = dbShowVideoList.Where("published_at >= ?", QueryWebVideoInfo.From)
	}
	if QueryWebVideoInfo.To != "" {
		dbShowVideoList = dbShowVideoList.Where("published_at <= ?", QueryWebVideoInfo.To)
	}

	log.Logger.Info(dbShowVideoList.Statement.SQL.String())
	err := dbShowVideoList.Find(&VideoList).Error
	if err != nil {
		log.Logger.Error(err.Error())
		return wholeResult
	} else {
		//for _, info := range TwoSetVideoInfos {
		//	logs.Info(info)
		//}
		wholeResult["total_num"] = len(VideoList)
		wholeResult["info_list"] = VideoList
		wholeResult["msg"] = "success"
		wholeResult["error_num"] = 0
	}

	//log.Logger.Info(fmt.Sprintf("ShowYoutubeVideoList resp------------%v", wholeResult))
	return wholeResult
}

package handles

import (
	"awesomeProject/app/serverCore/dto"
	"awesomeProject/middleware/db"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetKeyWord(c *gin.Context) {
	var keyWords []models.KeyWord
	var QueryKeyWord dto.QueryKeyWord
	bindJSONErr := c.ShouldBindJSON(&QueryKeyWord)
	if bindJSONErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindJSONErr.Error()})
		return
	}
	dbGetKeyWord := db.DB
	dbGetKeyWord = dbGetKeyWord.Order("update_time desc")
	dbGetKeyWord = dbGetKeyWord.Where("enable = ?", 1)
	dbGetKeyWord = dbGetKeyWord.Where("intellectual_property_names = ?", QueryKeyWord.IntellectualPropertyName)
	// 读取所有数据
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "获取萌点失败，请重试或者联系狗狗"}
	err := dbGetKeyWord.Debug().Find(&keyWords).Limit(2000).Error
	// Limit得放前面才生效
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wholeResult["total_num"] = len(keyWords)
	wholeResult["keyWord_list"] = keyWords
	wholeResult["msg"] = "success"
	wholeResult["error_num"] = 0
	c.JSON(http.StatusOK, wholeResult)
}

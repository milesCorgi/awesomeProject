package handles

import (
	"awesomeProject/middleware/db"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIntellectualPropertyNames(c *gin.Context) {
	var intellectualPropertyNames []models.IntellectualPropertyNames
	dbGetKeyWord := db.DB
	dbGetKeyWord = dbGetKeyWord.Order("update_time desc")
	dbGetKeyWord = dbGetKeyWord.Where("enable = ?", 1)
	// 读取所有数据
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "获取ip所属列表失败，请重试或者联系狗狗"}
	err := dbGetKeyWord.Debug().Find(&intellectualPropertyNames).Limit(2000).Error
	// Limit得放前面才生效
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wholeResult["total_num"] = len(intellectualPropertyNames)
	wholeResult["intellectual_property_name_list"] = intellectualPropertyNames
	wholeResult["msg"] = "success"
	wholeResult["error_num"] = 0
	c.JSON(http.StatusOK, wholeResult)
}

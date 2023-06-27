package handles

import (
	"awesomeProject/app/cache/dto"
	"awesomeProject/middleware/caches"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var cstSh, _ = time.LoadLocation("Asia/Shanghai")

func GetAllCache(c *gin.Context) {
	//var output map[string]caches.Element
	wholeResult := map[string]interface{}{"error_num": 0}
	//for k, v := range caches.GetAllCache() {
	//	fmt.Printf("key=%s, value=%v\n", k, v)
	//}
	wholeResult["caches"] = caches.GetAllCache()
	wholeResult["total_caches_num"] = len(caches.GetAllCache())
	c.JSON(http.StatusOK, wholeResult)
}

func ClearCache(c *gin.Context) {
	//var output map[string]caches.Element
	caches.ClearCache()
	c.JSON(http.StatusOK, map[string]interface{}{"error_num": 0, "msg": "success"})
}

func GetACache(c *gin.Context) {
	//var output map[string]caches.Element
	var cacheInfo dto.CacheInfo
	bindJSONErr := c.ShouldBindJSON(&cacheInfo)
	if bindJSONErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindJSONErr.Error()})
		return
	}
	if cacheInfo.Key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key can't be empty"})
		return
	}

	wholeResult := map[string]interface{}{"error_num": 0}
	wholeResult["key"] = cacheInfo.Key
	wholeResult["info"] = caches.GetCache(cacheInfo.Key)
	c.JSON(http.StatusOK, wholeResult)
}

func SetACache(c *gin.Context) {
	var cacheInfo dto.CacheInfo
	bindJSONErr := c.ShouldBindJSON(&cacheInfo)
	if bindJSONErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindJSONErr.Error()})
		return
	}
	if cacheInfo.Value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Value can't be empty"})
		return
	}
	if cacheInfo.Key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key can't be empty"})
		return
	}
	if cacheInfo.ExpiredSecond == "" {
		cacheInfo.ExpiredSecond = "5"
	}
	var element caches.Element
	element.Value = cacheInfo.Value
	now := time.Now()
	ExpiredSecondInt, err := strconv.ParseInt(cacheInfo.ExpiredSecond, 10, 64)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	element.ExpiredTime = now.Add(time.Duration(ExpiredSecondInt) * time.Second)
	caches.PutCache(cacheInfo.Key, element)
	wholeResult := map[string]interface{}{"error_num": 0, "msg": "success"}
	c.JSON(http.StatusOK, wholeResult)
}

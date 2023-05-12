package serverCore

import (
	"awesomeProject/middleware/db"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHander(c *gin.Context) {
	c.String(200, "pong")
}

func DBQueryDemo(c *gin.Context) {
	var videoNoteInfos []models.VideoNoteInfo
	err := db.DB.Find(&videoNoteInfos).Limit(1).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, videoNoteInfos)
}

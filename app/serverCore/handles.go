package serverCore

import (
	"github.com/gin-gonic/gin"
)

func PingHander(c *gin.Context) {
	c.String(200, "pong")
}

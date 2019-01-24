package topic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "topic/index", gin.H{
		"title":   "社区首页",
		"user_id": 1,
	})
}

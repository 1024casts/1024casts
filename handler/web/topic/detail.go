package topic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	c.HTML(http.StatusOK, "topic/detail", gin.H{
		"title":   "社区首页",
		"user_id": 1,
	})
}

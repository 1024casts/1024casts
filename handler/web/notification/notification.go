package notification

import (
	"net/http"

	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.HTML(http.StatusOK, "notification/list", gin.H{
		"title":   "通知",
		"user_id": util.GetUserId(c),
	})
}

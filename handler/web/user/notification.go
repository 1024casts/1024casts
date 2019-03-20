package user

import (
	"net/http"

	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Notification(c *gin.Context) {
	c.HTML(http.StatusOK, "user/notification", gin.H{
		"title":   "通知",
		"user_id": util.GetUserId(c),
	})
}

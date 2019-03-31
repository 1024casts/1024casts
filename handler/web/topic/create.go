package topic

import (
	"net/http"

	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	c.HTML(http.StatusOK, "user/follower", gin.H{
		"title":   "关注者",
		"user_id": util.GetUserId(c),
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

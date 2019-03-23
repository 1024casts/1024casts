package user

import (
	"net/http"

	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Following(c *gin.Context) {
	c.HTML(http.StatusOK, "user/following", gin.H{
		"title":   "正在关注",
		"user_id": util.GetUserId(c),
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

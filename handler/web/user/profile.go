package user

import (
	"net/http"

	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	c.HTML(http.StatusOK, "user/profile", gin.H{
		"title":     "个人资料",
		"user_id":   util.GetUserId(c),
		"user_name": util.GetUserId(c),
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

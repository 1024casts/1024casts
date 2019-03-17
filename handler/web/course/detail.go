package course

import (
	"net/http"

	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	slug := c.Param("slug")
	c.HTML(http.StatusOK, "course/detail", gin.H{
		"title":   "视频详情",
		"user_id": util.GetUserId(c),
		"slug":    slug,
	})
}

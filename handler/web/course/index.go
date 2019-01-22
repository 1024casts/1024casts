package course

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "course/index", gin.H{
		"title": "视频课程",
	})
}

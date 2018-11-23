package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "user/index", gin.H{
		"title": "布局页面",
	})
}

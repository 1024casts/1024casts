package wiki

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "wiki/index", gin.H{
		"title":   "wiki首页",
		"user_id": 1,
	})
}

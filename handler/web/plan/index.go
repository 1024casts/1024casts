package plan

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "plan/index", gin.H{
		"title":   "VIP订阅",
		"user_id": 1,
	})
}

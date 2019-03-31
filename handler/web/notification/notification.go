package notification

import (
	"net/http"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()

	user, err := srv.GetUserById(userId)
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404", gin.H{})
		return
	}
	c.HTML(http.StatusOK, "notification/list", gin.H{
		"title":   "通知",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
	})
}

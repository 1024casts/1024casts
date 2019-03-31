package topic

import (
	"net/http"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()

	user, _ := srv.GetUserById(userId)

	c.HTML(http.StatusOK, "topic/detail", gin.H{
		"title":   "话题详情",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
	})
}

package user

import (
	"net/http"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()

	c.HTML(http.StatusOK, "user/index", gin.H{
		"title":     "个人资料",
		"user_id":   userId,
		"user_name": srv.GetUserNameById(userId),
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

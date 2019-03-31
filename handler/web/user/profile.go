package user

import (
	"net/http"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()

	user, err := srv.GetUserById(userId)
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "user/profile", gin.H{
		"title":   "个人资料",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

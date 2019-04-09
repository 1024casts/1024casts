package user

import (
	"net/http"

	"github.com/1024casts/1024casts/service"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	userName := c.Param("username")
	srv := service.NewUserService()
	user, err := srv.GetUserByUsername(userName)
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404", gin.H{
			"title": "404错误",
			"ctx":   c,
		})
		return
	}

	c.HTML(http.StatusOK, "user/index", gin.H{
		"title":   "个人中心",
		"user_id": user.Id,
		"user":    user,
		"ctx":     c,
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

package user

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/1024casts/1024casts/pkg/pagination"

	"github.com/1024casts/1024casts/pkg/constvar"
	"github.com/lexkong/log"

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

	// get topic list
	topicSrv := service.NewTopicService()
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("get page error", err)
	}
	limit := constvar.DefaultLimit
	offset := (page - 1) * limit
	topicMap := make(map[string]interface{})
	topicMap["user_id"] = user.Id
	topics, count, err := topicSrv.GetTopicList(topicMap, offset, limit)
	if err != nil {
		log.Warnf("[topic] get topic list err: %v", err)
	}
	pagination := pagination.NewPagination(c.Request, count, limit)

	c.HTML(http.StatusOK, "user/index", gin.H{
		"title":   "个人中心",
		"user_id": user.Id,
		"user":    user,
		"ctx":     c,
		"topics":  topics,
		"pages":   template.HTML(pagination.Pages()),
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

package user

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/1024casts/1024casts/util"

	"github.com/1024casts/1024casts/pkg/pagination"

	"github.com/1024casts/1024casts/pkg/constvar"
	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/service"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()
	user, err := srv.GetUserById(userId)
	if err != nil {
		log.Warnf("[user] get user info err, %v", err)
	}

	userName := c.Param("username")
	userInfo, err := srv.GetUserByUsername(userName)
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404", gin.H{
			"title": "404错误",
			"ctx":   c,
		})
		return
	}

	// get topic list
	topicSrv := service.NewTopicService()
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Warnf("get page error", err)
	}
	limit := constvar.DefaultLimit
	offset := (page - 1) * limit
	topicMap := make(map[string]interface{})
	topicMap["user_id"] = userInfo.Id
	topics, count, err := topicSrv.GetTopicList(topicMap, offset, limit)
	if err != nil {
		log.Warnf("[topic] get topic list err: %v", err)
	}
	pagination := pagination.NewPagination(c.Request, count, limit)

	c.HTML(http.StatusOK, "user/index", gin.H{
		"title":    "个人中心",
		"user_id":  user.Id,
		"user":     user,
		"userInfo": userInfo,
		"ctx":      c,
		"topics":   topics,
		"pages":    template.HTML(pagination.Pages()),
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

package topic

import (
	"html/template"
	"net/http"

	"github.com/1024casts/1024casts/pkg/constvar"

	"github.com/1024casts/1024casts/pkg/pagination"

	"strconv"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Index(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewTopicService()

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("get page error", err)
	}
	limit := constvar.DefaultLimit
	offset := (page - 1) * limit
	topicMap := make(map[string]interface{})
	topics, count, err := srv.GetTopicList(topicMap, offset, limit)
	if err != nil {
		log.Warnf("[topic] get topic list err: %v", err)
	}

	pagination := pagination.NewPagination(c.Request, count, limit)
	c.HTML(http.StatusOK, "topic/index", gin.H{
		"title":   "社区首页",
		"user_id": userId,
		"ctx":     c,
		"topics":  topics,
		"pages":   template.HTML(pagination.Pages()),
	})
}

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

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
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

	userSrv := service.NewUserService()
	user, _ := userSrv.GetUserById(userId)

	pagination := pagination.NewPagination(c.Request, count, limit)

	//weekTopicMap := make(map[string]interface{})
	//weekTopicMap["created_at >="] = util.TimeToString(time.Now().Add(-time.Second * 86400 * 7))
	//weekTopics, count, err := srv.GetTopicList(weekTopicMap, offset, limit)
	//if err != nil {
	//	log.Warnf("[topic] get today topic list err: %v", err)
	//}

	// get top 15 by view_count
	topTopics, err := srv.GetTopTopicList(20)
	if err != nil {
		log.Warnf("[topic] get top topic list err: %+v", err)
	}

	c.HTML(http.StatusOK, "topic/index", gin.H{
		"title":     "社区首页",
		"user_id":   userId,
		"user":      user,
		"ctx":       c,
		"topics":    topics,
		"pages":     template.HTML(pagination.Pages()),
		"topTopics": topTopics,
	})
}

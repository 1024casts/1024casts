package topic

import (
	"net/http"
	"strconv"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	userId := util.GetUserId(c)
	userSrv := service.NewUserService()

	user, _ := userSrv.GetUserById(userId)

	topicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Warnf("[topic] get topic id err: %+v", err)
		topicId = 0
	}

	srv := service.NewTopicService()
	topic, err := srv.GetTopicById(topicId)
	if err != nil {
		log.Warnf("[topic] get topic detail err: %+v", err)
	}

	replyMap := make(map[string]interface{})
	replyMap["topic_id"] = topicId
	replies, _, err := srv.GetReplyList(replyMap, 0, 10)
	if err != nil {
		log.Warnf("[topic] get topic detail err: %+v", err)
	}

	// 增加view_count
	err = srv.IncrTopicViewCount(topicId)
	if err != nil {
		log.Warnf("[topic] incr topic view count err: %+v", err)
	}

	c.HTML(http.StatusOK, "topic/detail", gin.H{
		"title":   topic.Title,
		"user_id": userId,
		"user":    user,
		"topic":   topic,
		"replies": replies,
		"ctx":     c,
	})
}

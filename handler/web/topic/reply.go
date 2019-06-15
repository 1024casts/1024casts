package topic

import (
	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

type ReplyReq struct {
	TopicId    string `json:"topic_id" form:"topic_id"`
	OriginBody string `json:"origin_body" form:"origin_body"`
	Body       string `json:"body" form:"body"`
}

func Reply(c *gin.Context) {
	var req ReplyReq
	if err := c.Bind(&req); err != nil {
		app.Response(c, errno.ErrParam, nil)
		return
	}

	topicId := uint64(util.DecodeTopicId(req.TopicId))
	topicSrv := service.NewTopicService()
	_, err := topicSrv.GetTopicById(topicId)
	if err != nil {
		app.Response(c, errno.ErrDataIsNotExist, nil)
		return
	}

	replyModel := model.ReplyModel{
		TopicId:    topicId,
		OriginBody: req.OriginBody,
		Body:       req.Body,
		Source:     "PC",
		IsBlocked:  "no",
		UserId:     util.GetUserId(c),
	}

	userId := util.GetUserId(c)
	_, err = topicSrv.CreateReply(userId, topicId, replyModel)
	if err != nil {
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	app.Response(c, errno.OK, nil)
	return
}

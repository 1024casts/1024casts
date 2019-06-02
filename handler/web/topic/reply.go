package topic

import (
	"github.com/1024casts/1024casts/pkg/errno"

	"github.com/1024casts/1024casts/model"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

type ReplyReq struct {
	TopicId    int    `json:"topic_id" form:"topic_id"`
	OriginBody string `json:"origin_body" form:"origin_body"`
	Body       string `json:"body" form:"body"`
}

func Reply(c *gin.Context) {
	var req ReplyReq
	if err := c.Bind(&req); err != nil {
		app.Response(c, errno.ErrParam, nil)
		return
	}

	topicSrv := service.NewTopicService()
	replyModel := model.ReplyModel{
		TopicID:    req.TopicId,
		OriginBody: req.OriginBody,
		Body:       req.Body,
		Source:     "PC",
		IsBlocked:  "no",
		UserId:     util.GetUserId(c),
	}
	_, err := topicSrv.CreateReply(replyModel)
	if err != nil {
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	app.Response(c, errno.OK, nil)
	return
}

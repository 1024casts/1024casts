package topic

import (
	"strconv"

	"github.com/1024casts/1024casts/service"
	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Like(c *gin.Context) {
	replyId := c.Param("reply_id")
	rId, err := strconv.Atoi(replyId)
	if err != nil {
		app.Response(c, errno.ErrParam, nil)
		return
	}

	topicSrv := service.NewTopicService()
	reply, err := topicSrv.GetReplyById(rId)
	if err != nil {
		log.Warnf("[comment] get comment err: %v", err)
		app.Response(c, err, nil)
	}
	if reply.Id == 0 {
		app.Response(c, errno.ErrDataIsNotExist, nil)
		return
	}

	err = topicSrv.IncrReplyLikeCount(rId)
	if err != nil {
		log.Warnf("[topic] incr reply like count err: %v", err)
	}

	app.Response(c, nil, nil)
}

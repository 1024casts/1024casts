package comment

import (
	"fmt"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/pkg/notification"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

type CreateCommentReq struct {
	Type          int    `json:"type" form:"type"`
	RelatedId     int    `json:"related_id" form:"related_id"`
	OriginContent string `json:"origin_content" form:"origin_content"`
	Content       string `json:"content" form:"content"`
}

func Create(c *gin.Context) {

	var req CreateCommentReq
	if err := c.Bind(&req); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	cmtSrv := service.NewCommentService()
	commentModel := model.CommentModel{
		Type:          req.Type,
		RelatedId:     req.RelatedId,
		Ip:            "",
		Content:       util.MarkdownToHtml(util.ParseMentionUser(req.OriginContent)),
		OriginContent: req.OriginContent,
		LikeCount:     0,
		UserId:        util.GetUserId(c),
		DeviceType:    "",
	}
	cmtId, err := cmtSrv.CreateComment(commentModel)
	if err != nil {
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	// send to slack notification
	go func() {
		msg := fmt.Sprintf("new comment: %s[%d]", req.OriginContent, cmtId)
		err := notification.SendNewCommentSlackNotification(msg)
		if err != nil {
			log.Warnf("[comment] send new comment to slack err, %v", err)
		}
	}()

	app.Response(c, nil, cmtId)
}

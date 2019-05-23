package comment

import (
	"strconv"

	"github.com/1024casts/1024casts/service"
	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Like(c *gin.Context) {
	commentId := c.Param("comment_id")
	cmtId, err := strconv.Atoi(commentId)
	if err != nil {
		app.Response(c, errno.ErrParam, nil)
		return
	}

	cmtSrv := service.NewCommentService()
	comment, err := cmtSrv.GetCommentById(cmtId)
	if err != nil {
		log.Warnf("[comment] get comment err: %v", err)
		app.Response(c, err, nil)
	}
	if comment.Id == 0 {
		app.Response(c, errno.ErrDataIsNotExist, nil)
		return
	}

	cmtSrv.IncrLikeCount(cmtId)

	app.Response(c, nil, nil)
}

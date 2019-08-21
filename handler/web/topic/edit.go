package topic

import (
	"net/http"

	"github.com/1024casts/1024casts/pkg/errno"

	"github.com/1024casts/1024casts/model"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/service"
	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {

	userSrv := service.NewUserService()
	user, err := userSrv.GetUserById(util.GetUserId(c))
	if err != nil {
		log.Warnf("[topic] get user info err: %v", err)
	}

	topicSrv := service.NewTopicService()
	categories, err := topicSrv.GetCategoryList()
	if err != nil {
		log.Warnf("[topic] get categories err: %v", err)
	}

	topicIdStr := c.Param("id")
	topicId := util.DecodeTopicId(topicIdStr)
	topic, err := topicSrv.GetTopicById(uint64(topicId))
	if err != nil {
		log.Warnf("[topic] get topic info err: %v", err)
		app.Redirect(c, "/topics/"+topicIdStr, "参数错误")
		return
	}

	// check 用户是否有权限修改该topic
	if topic.User.Id != util.GetUserId(c) {
		log.Warnf("[topic] no have edit right for topic_id: %d", topic.Id)
		app.Redirect(c, "/topics/"+topicIdStr, errno.ErrNoRightEdit.Message)
		return
	}

	c.HTML(http.StatusOK, "topic/edit", gin.H{
		"title":      "编辑话题",
		"user_id":    util.GetUserId(c),
		"user":       user,
		"ctx":        c,
		"categories": categories,
		"topic":      topic,
	})
}

type EditTopicReq struct {
	Title      string `json:"title" form:"title"`
	CategoryId int    `json:"category_id" form:"category_id"`
	OriginBody string `json:"origin_body" form:"origin_body"`
	Body       string `json:"body" form:"body"`
}

func DoEdit(c *gin.Context) {
	var req EditTopicReq
	if err := c.Bind(&req); err != nil {
		app.Response(c, errno.ErrParam, nil)
		return
	}

	topicSrv := service.NewTopicService()
	topicId := util.DecodeTopicId(c.Param("id"))
	topic, err := topicSrv.GetTopicById(uint64(topicId))
	if err != nil {
		app.Response(c, errno.ErrDataIsNotExist, nil)
		return
	}

	// check 用户是否有权限修改该topic
	if topic.User.Id != util.GetUserId(c) {
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	topicModel := model.TopicModel{
		CategoryID: req.CategoryId,
		Title:      req.Title,
		OriginBody: req.OriginBody,
		Body:       util.MarkdownToHtml(req.OriginBody),
	}

	err = topicSrv.UpdateTopic(topicModel, uint64(topicId))
	if err != nil {
		app.Response(c, errno.ErrNoRightEdit, nil)
		return
	}

	app.Response(c, errno.OK, nil)
	return
}

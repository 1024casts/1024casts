package topic

import (
	"net/http"
	"strconv"

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
	topicId, err := strconv.Atoi(topicIdStr)
	if err != nil {
		log.Warnf("[topic] get topic id err: %v", err)
		app.Redirect(c, "/topics/"+topicIdStr, "参数错误")
		return
	}
	topic, err := topicSrv.GetTopicById(topicId)
	if err != nil {
		log.Warnf("[topic] get topic info err: %v", err)
		app.Redirect(c, "/topics/"+topicIdStr, "参数错误")
		return
	}

	c.HTML(http.StatusOK, "topic/edit", gin.H{
		"title":      "发布心话题",
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
	topicModel := model.TopicModel{
		CategoryID: req.CategoryId,
		Title:      req.Title,
		OriginBody: req.OriginBody,
		Body:       req.Body,
	}

	topicIdStr := c.Param("id")
	topicId, err := strconv.Atoi(topicIdStr)
	if err != nil {
		log.Warnf("[topic] get topic id err: %v", err)
		app.Response(c, errno.ErrParam, nil)
		return
	}

	err = topicSrv.UpdateTopic(topicModel, topicId)
	if err != nil {
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	app.Response(c, errno.OK, nil)
	return
}

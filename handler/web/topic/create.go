package topic

import (
	"net/http"

	"github.com/1024casts/1024casts/model"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/service"
	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

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

	c.HTML(http.StatusOK, "topic/create", gin.H{
		"title":      "发布心话题",
		"user_id":    util.GetUserId(c),
		"user":       user,
		"ctx":        c,
		"categories": categories,
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

type CreateTopicReq struct {
	Title      string `json:"title" form:"title"`
	CategoryId int    `json:"category_id" form:"category_id"`
	OriginBody string `json:"origin_body" form:"origin_body"`
	Body       string `json:"body" form:"body"`
}

func DoCreate(c *gin.Context) {
	var req CreateTopicReq
	if err := c.Bind(&req); err != nil {
		app.Redirect(c, "/topic/new", "参数错误")
		return
	}

	topicSrv := service.NewTopicService()
	topicModel := model.TopicModel{
		CategoryID:  req.CategoryId,
		Title:       req.Title,
		OriginBody:  req.OriginBody,
		Body:        req.Body,
		Source:      "PC",
		IsBlocked:   "no",
		IsExcellent: "no",
		UserID:      util.GetUserId(c),
	}
	_, err := topicSrv.CreateTopic(topicModel)
	if err != nil {
		app.Redirect(c, "/topic/new", "创建失败")
		return
	}

	app.Redirect(c, "/topics", "")
	return
}

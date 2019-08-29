package wiki

import (
	"net/http"
	"strconv"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Edit(c *gin.Context) {

	userSrv := service.NewUserService()
	user, err := userSrv.GetUserById(util.GetUserId(c))
	if err != nil {
		log.Warnf("[topic] get user info err: %v", err)
	}

	wikiSrv := service.NewWikiService()

	slug := c.Param("slug")
	wiki, err := wikiSrv.GetWikiPageBySlug(slug)
	if err != nil {
		log.Warnf("[wiki] get wiki info err: %v", err)
		app.Redirect(c, "/wiki/"+slug, errno.ErrNoRightEdit.Message)
		return
	}

	// check 用户是否有权限修改该wiki
	if wiki.UserId != util.GetUserId(c) {
		log.Warnf("[wiki] no have edit right for wiki_id: %d", wiki.Id)
		app.Redirect(c, "/wiki/"+slug, errno.ErrNoRightEdit.Message)
		return
	}

	c.HTML(http.StatusOK, "wiki/edit", gin.H{
		"title":   "编辑wiki",
		"user_id": util.GetUserId(c),
		"user":    user,
		"ctx":     c,
		"wiki":    wiki,
	})
}

type EditTopicReq struct {
	Title string `json:"title" form:"title"`
	//CategoryId int    `json:"category_id" form:"category_id"`
	OriginContent string `json:"origin_content" form:"origin_content"`
}

func DoEdit(c *gin.Context) {
	var req EditTopicReq
	if err := c.Bind(&req); err != nil {
		app.Response(c, errno.ErrParam, nil)
		return
	}

	wikiSrv := service.NewWikiService()
	wikiIdStr := c.Param("id")
	wikiId, _ := strconv.Atoi(wikiIdStr)
	log.Infof("wiki id: %v", wikiId)
	wiki, err := wikiSrv.GetWikiById(wikiId)
	if err != nil {
		app.Response(c, errno.ErrDataIsNotExist, nil)
		return
	}

	// check 用户是否有权限修改该topic
	if wiki.UserId != util.GetUserId(c) {
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	wikiPageModel := model.WikiPageModel{
		Title:         req.Title,
		OriginContent: req.OriginContent,
		Content:       util.MarkdownToHtml(req.OriginContent),
		FixCount:      wiki.FixCount + 1,
	}

	err = wikiSrv.UpdateWiki(int(wiki.Id), wikiPageModel)
	if err != nil {
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	app.Response(c, errno.OK, nil)
	return
}

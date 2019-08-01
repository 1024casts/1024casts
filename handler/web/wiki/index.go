package wiki

import (
	"net/http"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	userId := util.GetUserId(c)
	log.Infof("[wiki] get user id from util, uid: %v", userId)
	srv := service.NewUserService()
	user, err := srv.GetUserById(userId)
	if err != nil {
		log.Warnf("[wiki] get user info err, %v", err)
	}

	wikiSrv := service.NewWikiService()
	homeSlug := "index"
	wiki, err := wikiSrv.GetWikiPageBySlug(homeSlug)
	if err != nil {
		log.Warnf("[wiki] get wiki page info err: %v", err)
	}

	categories, err := wikiSrv.GetWikiCategoryListWithPage()
	if err != nil {
		log.Warnf("[wiki] get category with pages err: %v", err)
	}

	c.HTML(http.StatusOK, "wiki/index", gin.H{
		"title":      "wiki首页",
		"user_id":    userId,
		"user":       user,
		"ctx":        c,
		"categories": categories,
		"wiki":       wiki,
	})
}

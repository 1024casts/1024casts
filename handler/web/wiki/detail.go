package wiki

import (
	"net/http"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()
	user, _ := srv.GetUserById(userId)

	wikiSrv := service.NewWikiService()
	wiki, err := wikiSrv.GetWikiBySlug(c.Param("slug"))
	if err != nil {
		log.Warnf("[wiki] get wiki page info err: %v", err)
	}

	c.HTML(http.StatusOK, "wiki/detail", gin.H{
		"title":   wiki.Title,
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"wiki":    wiki,
	})
}

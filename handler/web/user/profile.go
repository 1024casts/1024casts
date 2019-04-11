package user

import (
	"net/http"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Profile(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()

	user, err := srv.GetUserById(userId)
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "user/profile", gin.H{
		"title":   "个人资料",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

type ProfileRequest struct {
	City            string `json:"city" form:"city"`
	Company         string `json:"company" form:"company"`
	GithubId        string `json:"github_id" form:"github_id"`
	PersonalWebsite string `json:"personal_website" form:"personal_website"`
}

// 更新基本资料
func DoProfile(c *gin.Context) {
	var req ProfileRequest
	if err := c.Bind(&req); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	userId := util.GetUserId(c)
	srv := service.NewUserService()

	log.Warnf("basic, info: %v", c.Request.RequestURI)

	_, err := srv.GetUserById(userId)
	if err != nil {
		app.Response(c, errno.ErrUserNotFound, nil)
		return
	}

	userMap := map[string]interface{}{
		"city":             req.City,
		"company":          req.Company,
		"github_id":        req.GithubId,
		"personal_website": req.PersonalWebsite,
	}
	err = srv.UpdateUser(userMap, userId)
	if err != nil {
		log.Warnf("[user] update user profile err: %v", err)
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	app.Response(c, errno.OK, nil)
	return
}

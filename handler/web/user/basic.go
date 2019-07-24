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

func Basic(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()

	log.Warnf("basic, info: %v", c.Request.RequestURI)

	user, err := srv.GetUserById(userId)
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "user/basic", gin.H{
		"title":   "个人资料",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

type BasicRequest struct {
	RealName     string `json:"real_name" form:"real_name"`
	Introduction string `json:"introduction" form:"introduction"`
	Avatar       string `json:"avatar" form:"avatar"`
}

// 更新基本资料
func DoBasic(c *gin.Context) {
	// Binding the course data.
	var req BasicRequest
	if err := c.Bind(&req); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	userId := util.GetUserId(c)
	srv := service.NewUserService()

	log.Warnf("[basic] info: %v", c.Request.RequestURI)

	_, err := srv.GetUserById(userId)
	if err != nil {
		app.Response(c, errno.ErrUserNotFound, nil)
		return
	}

	// single file
	avatar, err := c.FormFile("avatar")
	if err != nil {
		log.Warnf("[basic] get avatar err: %v", err)
		app.Response(c, errno.ErrGetUploadFile, nil)
		return
	}

	qiNiuSrv := service.NewQiNiuService()
	uploadRet, err := qiNiuSrv.UploadImage(c, avatar, false)
	if err != nil {
		log.Warnf("[basic] upload avatar err: %v", err)
		app.Response(c, errno.ErrUploadingFile, nil)
		return
	}

	userMap := map[string]interface{}{
		"real_name":    req.RealName,
		"introduction": req.Introduction,
		"avatar":       uploadRet.Key,
	}
	err = srv.UpdateUser(userMap, userId)
	if err != nil {
		log.Warnf("[basic] update user is_activation err: %v", err)
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	app.Response(c, errno.OK, service.UploadResponse{Key: uploadRet.Key, Url: uploadRet.Url})
	return
}

package user

import (
	"net/http"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/auth"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Password(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()

	user, err := srv.GetUserById(userId)
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "user/password", gin.H{
		"title":   "修改密码",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

type PasswordRequest struct {
	CurrentPassword      string `json:"current_password" form:"current_password"`
	Password             string `json:"password" form:"password"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation"`
}

// 更新基本资料
func DoPassword(c *gin.Context) {
	var req PasswordRequest
	if err := c.Bind(&req); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	userId := util.GetUserId(c)
	srv := service.NewUserService()

	user, err := srv.GetUserById(userId)
	if err != nil {
		app.Response(c, errno.ErrUserNotFound, nil)
		return
	}

	// check 旧密码是否和数据库密码一致
	if err := auth.Compare(user.Password, req.CurrentPassword); err != nil {
		log.Warn("[user] current password is neq password")
		app.Response(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// check 新密码是否和确认密码一致
	if req.Password != req.PasswordConfirmation {
		log.Warnf("[user] password neq user password_confirmation err: %v", err)
		app.Response(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	hashedPassword, err := auth.Encrypt(req.Password)
	if err != nil {
		log.Warnf("[user] Encrypt user password err: %v", err)
		app.Response(c, errno.ErrEncrypt, nil)
		return
	}

	userMap := map[string]interface{}{
		"password": hashedPassword,
	}
	err = srv.UpdateUser(userMap, userId)
	if err != nil {
		log.Warnf("[user] update user password err: %v", err)
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	// 重新登录
	util.ClearLoginCookie(c)

	app.Response(c, errno.OK, nil)
	return
}

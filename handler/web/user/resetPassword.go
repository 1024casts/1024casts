package user

import (
	"net/http"

	"fmt"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/auth"
	"github.com/1024casts/1024casts/service"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func ShowResetForm(c *gin.Context) {
	c.HTML(http.StatusOK, "user/resetPassword", gin.H{
		"title":   "重置密码",
		"user_id": 0,
		"ctx":     c,
		"token":   c.Param("token"),
	})
}

type ResetRequest struct {
	Email                string `json:"email" form:"email"`
	Password             string `json:"password" form:"password"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation"`
	Token                string `json:"token" form:"token"`
}

// 重置密码
func ResetPassword(c *gin.Context) {
	var req ResetRequest
	if err := c.Bind(&req); err != nil {
		app.Redirect(c, "/password/reset", "重置密码链接错误")
		return
	}
	var redirectPath = fmt.Sprintf("/password/reset/%s", req.Token)

	srv := service.NewUserService()

	// step1: check 新密码是否和确认密码一致
	if req.Password != req.PasswordConfirmation {
		log.Warnf("[user] password neq user password_confirmation err")
		app.Redirect(c, redirectPath, "两次密码输入不一致")
		return
	}

	// step2: check输入的email是否存在
	user, err := srv.GetUserByEmail(req.Email)
	if err != nil {
		log.Warnf("[user] get user info err by email: %s, err: %v", req.Email, err)
		app.Redirect(c, redirectPath, "Email输入错误，请检查")
		return
	}

	// step3: 根据email获取token
	hashedToken, err := srv.GetResetPasswordTokenByEmail(req.Email)
	if err != nil {
		log.Warnf("[user] get reset password token err: %v", err)
		app.Redirect(c, redirectPath, "输入的Email有误")
		return
	}

	// step4: 判断token与form传入的是否match
	if err := auth.Compare(hashedToken, req.Token); err != nil {
		log.Warnf("[user] current reset link invalid")
		app.Redirect(c, redirectPath, "当前重置链接无效")
		return
	}

	// step5: 更新数据库密码
	err = srv.UpdateUserPassword(user.Id, req.Password)
	if err != nil {
		log.Warnf("[user] update user password err: %v", err)
		app.Redirect(c, redirectPath, "重置出错")
		return
	}

	// step6: 删除reset表中的记录
	err = srv.DeleteResetPasswordByEmail(req.Email)
	if err != nil {
		log.Warnf("[user] update user password err: %v", err)
		app.Redirect(c, redirectPath, "重置出错")
		return
	}

	// 重新登录
	srv.ClearLoginCookie(c, user.Id)

	app.Redirect(c, redirectPath, "密码重置成功")
	return
}

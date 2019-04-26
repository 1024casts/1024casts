package user

import (
	"net/http"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/auth"
	"github.com/1024casts/1024casts/pkg/mail"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func ShowLinkRequestForm(c *gin.Context) {
	c.HTML(http.StatusOK, "user/forgetPassword", gin.H{
		"title":   "重置密码",
		"user_id": 0,
		"ctx":     c,
	})
}

type SendResetLinkRequest struct {
	Email string `json:"email" form:"email"`
}

// 发送重置密码邮件
func SendResetLinkEmail(c *gin.Context) {
	var redirectPath = "/password/reset"

	var req SendResetLinkRequest
	if err := c.Bind(&req); err != nil {
		app.Redirect(c, redirectPath, "邮箱地址填写错误")
		return
	}

	// step1: 校验email
	if req.Email == "" || !com.IsEmail(req.Email) {
		log.Warnf("[forget] email addr err")
		app.Redirect(c, redirectPath, "错误的邮件地址~")
		return
	}

	// step2: 根据email 获取用户信息
	srv := service.NewUserService()
	user, err := srv.GetUserByEmail(req.Email)
	if err != nil {
		app.Redirect(c, redirectPath, "错误的邮件地址")
		return
	}

	// ste3: 生成激活token to db
	token := util.GenPasswordToken()
	hashedToken, err := auth.Encrypt(token)
	if err != nil {
		log.Warnf("[user] encrypt token err:%v", err)
		app.Redirect(c, redirectPath, "重置发生错误")
		return
	}
	pwdReset := model.PasswordResetModel{
		Email: user.Email,
		Token: hashedToken,
	}
	model.DB.Self.Create(&pwdReset)

	// step4: 给指定用户发送邮件
	mailer := mail.NewMailer()
	go mailer.SendForgetPasswordMail(user.Email, token)

	app.Redirect(c, redirectPath, "我们已经发送密码重置链接到您的邮箱")
	return
}

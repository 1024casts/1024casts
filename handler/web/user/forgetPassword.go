package user

import (
	"bytes"
	"html/template"
	"net/http"

	"fmt"

	"io/ioutil"

	"time"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
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
	token, err := util.GenShortId()
	if err != nil {
		log.Warnf("[user] gen user reset token err: %v", err)
		app.Redirect(c, redirectPath, "内部错误")
		return
	}
	// write email and token do db

	// step4: 给指定用户发送邮件
	go sendResetMail(user.Username, user.Email, token)

	app.Redirect(c, redirectPath, "我们已经发送密码重置链接到您的邮箱")
	return
}

type mailData struct {
	HomeUrl       string `json:"home_url"`
	WebsiteName   string `json:"website_name"`
	WebsiteDomain string `json:"website_domain"`
	ResetUrl      string `json:"reset_url"`
	Year          int    `json:"year"`
}

// 发送密码重置邮件
func sendResetMail(username, toMail, token string) {
	m := gomail.NewMessage()
	// 发件人
	m.SetAddressHeader("From", "no-reply@phpcasts.org", viper.GetString("website.name"))
	// 收件人
	m.SetHeader("To",
		m.FormatAddress(toMail, ""),
	)
	// 主题
	m.SetHeader("Subject", "密码重置")
	// 正文
	resetUrl := fmt.Sprintf("%s/password/reset/%s", viper.GetString("website.domain"), token)

	mailTpl := readEmailTplContent()
	tpl, err := template.New("reset password tpl").Parse(mailTpl)
	if err != nil {
		log.Warnf("[user] template new err: %v", err)
		return
	}
	mailData := mailData{
		HomeUrl:       viper.GetString("website.domain"),
		WebsiteName:   viper.GetString("website.name"),
		WebsiteDomain: viper.GetString("website.domain"),
		ResetUrl:      resetUrl,
		Year:          time.Now().Year(),
	}
	buffer := new(bytes.Buffer)
	tpl.Execute(buffer, mailData)

	m.SetBody("text/html", buffer.String())
	//m.SetBody("text/html", "Hi, "+username+"<br>您的密码重置链接为： <a href = '"+resetUrl+"'>"+resetUrl+"</a>")

	// 发送邮件服务器、端口、发件人账号、发件人密码
	d := gomail.NewDialer(viper.GetString("mail.host"), viper.GetInt("mail.port"), viper.GetString("mail.username"), viper.GetString("mail.password"))
	if err := d.DialAndSend(m); err != nil {
		log.Warnf("[user] send reset mail err: %v", err)
	}
}

func readEmailTplContent() string {
	b, err := ioutil.ReadFile("./templates/email/reset-mail.html")
	if err != nil {
		log.Warnf("[user] read file err: %v", err)
		return ""
	}
	str := string(b)
	return str
}

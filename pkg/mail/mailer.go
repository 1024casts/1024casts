package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"time"

	"io/ioutil"

	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	Subject  string `json:"subject"`
	ToMail   string `json:"to_mail"`
	HtmlBody string `json:"html_body"`
}

func NewMailer() *Mailer {
	return &Mailer{}
}

type ForgetPasswordMailData struct {
	HomeUrl       string `json:"home_url"`
	WebsiteName   string `json:"website_name"`
	WebsiteDomain string `json:"website_domain"`
	ResetUrl      string `json:"reset_url"`
	Year          int    `json:"year"`
}

func (mail *Mailer) SendForgetPasswordMail(toMail, token string) {
	mail.Subject = "密码重置"
	mail.ToMail = toMail
	// 正文
	resetUrl := fmt.Sprintf("%s/password/reset/%s", viper.GetString("website.domain"), token)
	mailData := ForgetPasswordMailData{
		HomeUrl:       viper.GetString("website.domain"),
		WebsiteName:   viper.GetString("website.name"),
		WebsiteDomain: viper.GetString("website.domain"),
		ResetUrl:      resetUrl,
		Year:          time.Now().Year(),
	}
	mailTplContent := mail.getEmailTplContent("./templates/email/reset-mail.html", mailData)
	mail.HtmlBody = mailTplContent

	mail.sendMail()
}

func (mail *Mailer) getEmailTplContent(tplPath string, mailData interface{}) string {
	b, err := ioutil.ReadFile(tplPath)
	if err != nil {
		log.Warnf("[user] read file err: %v", err)
		return ""
	}
	mailTpl := string(b)
	tpl, err := template.New("reset password tpl").Parse(mailTpl)
	if err != nil {
		log.Warnf("[user] template new err: %v", err)
		return ""
	}
	buffer := new(bytes.Buffer)
	tpl.Execute(buffer, mailData)
	return buffer.String()
}

func (mail *Mailer) sendMail() {
	m := gomail.NewMessage()
	// 发件人
	m.SetAddressHeader("From", "no-reply@phpcasts.org", viper.GetString("website.name"))
	// 收件人
	m.SetHeader("To",
		m.FormatAddress(mail.ToMail, ""),
	)
	// 主题
	m.SetHeader("Subject", mail.Subject)

	m.SetBody("text/html", mail.HtmlBody)
	//m.SetBody("text/html", "Hi, "+username+"<br>您的密码重置链接为： <a href = '"+resetUrl+"'>"+resetUrl+"</a>")

	// 发送邮件服务器、端口、发件人账号、发件人密码
	d := gomail.NewDialer(viper.GetString("mail.host"), viper.GetInt("mail.port"), viper.GetString("mail.username"), viper.GetString("mail.password"))
	if err := d.DialAndSend(m); err != nil {
		log.Warnf("[user] send reset mail err: %v", err)
	}
}

package user

import (
	"net/http"

	"github.com/spf13/viper"

	. "github.com/1024casts/1024casts/handler"

	"github.com/gorilla/sessions"
	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {

	log.Infof("[user] entry logout...")

	// 删除cookie信息
	session := GetCookieSession(c)
	session.Options = &sessions.Options{
		Domain: viper.GetString("cookie.domain"),
		Path:   "/",
		MaxAge: -1,
	}
	err := session.Save(Request(c), ResponseWriter(c))
	if err != nil {
		log.Warnf("[user] logout save session err: %v", err)
		c.Abort()
		return
	}

	// 重定向得到原页面
	c.Redirect(http.StatusSeeOther, c.Request.Referer())
	return
}

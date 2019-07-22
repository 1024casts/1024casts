package handler

import (
	"net/http"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

var Store = sessions.NewCookieStore([]byte("cookie-secret-1024casts"))

func SetLoginCookie(ctx *gin.Context, userId uint64) {
	Store.Options.HttpOnly = true

	session := GetCookieSession(ctx)
	session.Options = &sessions.Options{
		Domain:   viper.GetString("cookie.domain"),
		MaxAge:   86400,
		Path:     "/",
		HttpOnly: true,
	}
	// 浏览器关闭，cookie删除，否则保存30天(github.com/gorilla/sessions 包的默认值)
	val, ok := ctx.GetPostForm("remember_me")
	log.Infof("[handler] remember_me val: %s, ok: %v", val, ok)
	if ok && val == "1" {
		session.Options.MaxAge = 86400 * 30
	}

	session.Values["user_id"] = userId

	req := Request(ctx)
	resp := ResponseWriter(ctx)
	err := session.Save(req, resp)
	if err != nil {
		log.Warnf("[handler] set login cookie, %v", err)
	}
}

func GetCookieSession(ctx *gin.Context) *sessions.Session {
	session, err := Store.Get(ctx.Request, viper.GetString("cookie.name"))
	if err != nil {
		log.Warnf("[handler] store get err, %v", err)
	}
	return session
}

func Request(ctx *gin.Context) *http.Request {
	return ctx.Request
}

func ResponseWriter(ctx *gin.Context) http.ResponseWriter {
	return ctx.Writer
}

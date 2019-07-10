package middleware

import (
	"net/http"

	"github.com/1024casts/1024casts/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

// see: http://researchlab.github.io/2016/03/29/gin-setcookie/
func CookieMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(viper.GetString("cookie.name"))
		log.Infof("current cookie: %s", cookie)
		if err != nil {
			log.Warnf("[middleware] err: %v", err)
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		}

		if len(cookie) == 0 {
			log.Warnf("[middleware] current cookie len is zero")
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
		} else {
			userId := util.GetUserId(c)
			log.Infof("current user_id: %d from cookie", userId)
			util.SetLoginCookie(c, userId)
		}

		c.Next()
	}
}

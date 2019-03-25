package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

func CookieMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie(viper.GetString("cookie.name"))
		log.Infof("current cookie: %s", cookie)
		if len(cookie) == 0 {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
		} else {
			c.SetCookie(viper.GetString("cookie.name"), cookie, viper.GetInt("cookie.max_age"), "/", "localhost:8888", false, true)
		}

		c.Next()
	}
}

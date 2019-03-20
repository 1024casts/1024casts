package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

func CookieMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie(viper.GetString("cookie.name"))
		if err != nil {
			log.Warnf("[middleware] get cookie err: %+v", err)
			c.Redirect(http.StatusMovedPermanently, "/login")
		}
		value := cookie.Value

		c.Next()

		if len(value) == 0 {
			cookie.Value = ""
			log.Infof("current cookie %+v", cookie)
			c.Redirect(http.StatusMovedPermanently, "/login")
		} else {
			c.SetCookie(viper.GetString("cookie.name"), cookie.Value, viper.GetInt("cookie.max_age"), "/", "localhost:8888", false, true)
		}
	}
}

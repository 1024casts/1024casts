package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func CookieMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.SetCookie("Cookie", "123456", 3600, "/", "localhost:8888", false, true)

		if cookie, err := c.Request.Cookie("Cookie"); err == nil {
			value := cookie.Value

			c.Next()

			if len(value) == 0 {
				cookie.Value = ""
				log.Infof("cookie...... %+v", cookie)
				c.Redirect(http.StatusMovedPermanently, "/login")
			} else {
				c.SetCookie("Cookie", cookie.Value, 3600, "/", "localhost:8888", false, true)
			}
		} else {
			log.Warnf("Get cookie err")
		}
	}
}

package middleware

import (
	"net/http"

	"github.com/1024casts/1024casts/handler"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// see: http://researchlab.github.io/2016/03/29/gin-setcookie/
func CookieMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := handler.GetCookieSession(c)
		log.Infof("current session: %s", session)
		userId, ok := session.Values["user_id"]
		log.Infof("current user_id: %d from cookie", userId)
		if ok {
			handler.SetLoginCookie(c, userId.(uint64))
		} else {
			log.Warnf("[middleware] current user_id is not ok")
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		}

		c.Next()
	}
}

package middleware

import "github.com/gin-gonic/gin"

// Gin 框架官方推荐了一款中间件
// more: https://github.com/aviddiviner/gin-limit

func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()

	}
}

package app

import (
	"net/http"

	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/pkg/flash"
	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Resp{
		Code:    code,
		Message: message,
		Data:    data,
	})

	return
}

func Redirect(c *gin.Context, redirectPath, errMsg string) {
	flash.SetFlash(c.Writer, "error", []byte(errMsg))
	c.Redirect(http.StatusMovedPermanently, redirectPath)
	c.Abort()
}

package notification

import (
	"net/http"

	"html/template"

	"github.com/1024casts/1024casts/pkg/pagination"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()

	user, err := srv.GetUserById(userId)
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404", gin.H{})
		return
	}

	//创建一个分页器，一万条数据，每页30条
	pagination := pagination.NewPagination(c.Request, 10000, 30)
	//传到模板中需要转换成template.HTML类型，否则html代码会被转义
	c.HTML(http.StatusOK, "notification/list", gin.H{
		"title":   "通知",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"pages":   template.HTML(pagination.Pages()),
	})
}

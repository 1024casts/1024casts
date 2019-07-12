package course

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/1024casts/1024casts/pkg/constvar"
	"github.com/1024casts/1024casts/pkg/pagination"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()
	user, _ := srv.GetUserById(userId)

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Error("get page error", err)
	}
	limit := constvar.DefaultLimit
	offset := (page - 1) * limit

	courseSrv := service.NewCourseService()
	courseMap := make(map[string]interface{})
	courseMap["is_publish"] = 1
	courses, count, err := courseSrv.GetCourseList(courseMap, offset, 12)
	if err != nil {
		log.Warnf("[course] get course list err: %v", err)
	}

	pagination := pagination.NewPagination(c.Request, count, limit)
	c.HTML(http.StatusOK, "course/index", gin.H{
		"title":   "视频课程",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"courses": courses,
		"pages":   template.HTML(pagination.Pages()),
	})
}

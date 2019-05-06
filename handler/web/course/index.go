package course

import (
	"net/http"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()
	user, _ := srv.GetUserById(userId)

	courseSrv := service.NewCourseService()
	courseMap := make(map[string]interface{})
	courseMap["is_publish"] = 1
	courses, count, err := courseSrv.GetCourseList(courseMap, 0, 20)
	if err != nil {
		log.Warnf("[course] get course list err: %v", err)
	}

	c.HTML(http.StatusOK, "course/index", gin.H{
		"title":   "视频课程",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"count":   count,
		"courses": courses,
	})
}

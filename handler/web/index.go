package web

import (
	"net/http"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
	log "qiniupkg.com/x/log.v7"
)

func Index(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()

	user, _ := srv.GetUserById(userId)

	courseSrv := service.NewCourseService()
	courseMap := make(map[string]interface{})
	courseMap["is_publish"] = 1
	courses, _, err := courseSrv.GetCourseList(courseMap, 0, 6)
	if err != nil {
		log.Warnf("[course] get course list err: %v", err)
	}

	c.HTML(http.StatusOK, "index", gin.H{
		"title":   "首页",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"courses": courses,
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

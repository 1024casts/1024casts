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

	// get course list
	courseMap := make(map[string]interface{})
	courseMap["is_publish"] = 1
	courses, _, err := courseSrv.GetCourseList(courseMap, 0, 8)
	if err != nil {
		log.Warnf("[index] get course list err: %v", err)
	}

	// get course total count
	courseCount, err := courseSrv.GetCourseTotalCount(courseMap)
	if err != nil {
		log.Warnf("[index] get course total count err, %v", err)
	}

	// get video total count
	videoSrv := service.NewVideoService()
	videoCount, err := videoSrv.GetVideoTotalCount()
	if err != nil {
		log.Warnf("[index] get video total count err, %v", err)
	}

	// get video total minute(总分钟数)
	videoDuration, err := videoSrv.GetVideoTotalDuration()
	if err != nil {
		log.Warnf("[index] get video total duration err, %v", err)
	}
	// 视频总分钟数
	videoTotalMinute := videoDuration / 60

	c.HTML(http.StatusOK, "index", gin.H{
		"title":            "首页",
		"user_id":          userId,
		"user":             user,
		"ctx":              c,
		"courses":          courses,
		"courseTotalCount": courseCount,
		"videoTotalCount":  videoCount,
		"videoTotalMinute": videoTotalMinute,
	})
}

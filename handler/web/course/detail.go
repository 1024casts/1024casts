package course

import (
	"net/http"

	"github.com/1024casts/1024casts/service"
	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	slug := c.Param("slug")

	courseSrv := service.NewCourseService()
	course, err := courseSrv.GetCourseBySlug(slug)
	if err != nil {
		log.Warnf("[course] get course list err: %v", err)
	}

	sections, err := courseSrv.GetCourseSectionList(course.Id)
	if err != nil {
		log.Warnf("[course] get video list err: %+v", err)
	}

	c.HTML(http.StatusOK, "course/detail", gin.H{
		"title":    "视频详情",
		"user_id":  util.GetUserId(c),
		"slug":     slug,
		"ctx":      c,
		"course":   course,
		"sections": sections,
	})
}

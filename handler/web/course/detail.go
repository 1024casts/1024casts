package course

import (
	"net/http"

	"github.com/1024casts/1024casts/model"

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

	// 视频展示是否分组
	isGroup := false
	videoSrv := service.NewVideoService()
	videos, err := videoSrv.GetVideoList(course.Id, false)
	if err != nil {
		log.Warnf("[course] get video list err: %v", err)
	}

	sections := make([]*model.SectionModel, 0)
	// 无视频说明进行了分组
	//if len(videos) == 0 {
	//	isGroup = true
	sections, err = courseSrv.GetCourseSectionListWithVideo(course.Id)
	if err != nil {
		log.Warnf("[course] get video list err: %+v", err)
	}
	//}
	if len(sections) > 0 {
		isGroup = true
	}

	c.HTML(http.StatusOK, "course/detail", gin.H{
		"title":    "视频详情",
		"user_id":  util.GetUserId(c),
		"slug":     slug,
		"ctx":      c,
		"course":   course,
		"isGroup":  isGroup,
		"sections": sections,
		"videos":   videos,
	})
}

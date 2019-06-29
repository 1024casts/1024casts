package video

import (
	"net/http"
	"strconv"

	"github.com/1024casts/1024casts/pkg/constvar"

	"github.com/1024casts/1024casts/service"
	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	slug := c.Param("slug")

	userSrv := service.NewUserService()
	user, err := userSrv.GetUserById(util.GetUserId(c))
	if err != nil {
		log.Warnf("[video] get course info err: %v", err)
	}

	courseSrv := service.NewCourseService()
	course, err := courseSrv.GetCourseBySlug(slug)
	if err != nil {
		log.Warnf("[video] get course info err: %v", err)
	}

	episodeId, err := strconv.Atoi(c.Param("episode_id"))
	if err != nil {
		log.Warnf("[video] get param episode_id err: %v", err)
	}

	videoSrv := service.NewVideoService()
	video, err := videoSrv.GetVideoByCourseIdAndEpisodeId(course.Id, episodeId)
	if err != nil {
		log.Warnf("[video] get video info err: %+v", err)
	}

	recentCourses, err := courseSrv.GetRecentCourses(15)
	if err != nil {
		log.Warnf("[video] get recent course err: %+v", err)
	}

	commentSrv := service.NewCommentService()
	commentMap := make(map[string]interface{})
	commentMap["type"] = constvar.CommentTypeVideo
	commentMap["related_id"] = video.Id
	comments, cmtCount, err := commentSrv.GetCommentList(commentMap, 0, 1000)
	if err != nil {
		log.Warnf("[video] get comments err: %+v", err)
		return
	}

	c.HTML(http.StatusOK, "video/detail", gin.H{
		"title":         video.Name + " - 视频详情",
		"user_id":       util.GetUserId(c),
		"user":          user,
		"slug":          slug,
		"ctx":           c,
		"course":        course,
		"video":         video,
		"recentCourses": recentCourses,
		"cmtCount":      cmtCount,
		"comments":      comments,
	})
}

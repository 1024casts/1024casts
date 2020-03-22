package video

import (
	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary Add new video to the database
// @Description Add a new course
// @Tags course
// @Accept  json
// @Produce  json
// @Param course body course.CreateRequest true "Create a new course"
// @Success 200 {object} course.CreateResponse "{"code":0,"message":"OK","data":{"name":"test"}}"
// @Router /videos [post]
func Create(c *gin.Context) {
	log.Info("Video Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	item := model.VideoModel{
		CourseID:    r.CourseID,
		SectionID:   r.SectionID,
		EpisodeID:   r.EpisodeID,
		Name:        r.Name,
		Description: r.Description,
		Mp4Key:      r.Mp4Key,
		Duration:    r.Duration,
		CoverKey:    r.CoverKey,
		IsFree:      r.IsFree,
		IsPublish:   r.IsPublish,
	}

	srv := service.NewVideoService()
	id, err := srv.CreateVideo(item)
	if err != nil {
		app.Response(c, errno.ErrVideoCreateFail, nil)
		log.Warn("video info", lager.Data{"id": id})
		return
	}

	resp := CreateResponse{
		Id: item.Id,
	}

	// Show the user information.
	app.Response(c, nil, resp)
}

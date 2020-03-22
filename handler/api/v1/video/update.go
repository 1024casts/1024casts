package video

import (
	"strconv"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/util"

	"github.com/1024casts/1024casts/service"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary Update a course info by the course identifier
// @Description Update a course by ID
// @Tags course
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The course's database id index num"
// @Param user body model.CourseModel true "The course info"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /videos/{id} [put]
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// Get the course id from the url parameter.
	videoId, _ := strconv.Atoi(c.Param("id"))

	// Binding the course data.
	var video CreateRequest
	if err := c.Bind(&video); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	srv := service.NewVideoService()
	//user, err := model.GetUserById(userId)
	_, err := srv.GetVideoById(videoId)
	if err != nil {
		app.Response(c, errno.ErrVideoNotFound, nil)
		log.Warn("video info", lager.Data{"id": videoId})
		return
	}

	// Validate the data.
	//if err := u.Validate(); err != nil {
	//	app.Response(c, errno.ErrValidation, nil)
	//	return
	//}

	// Save changed fields.
	itemMap := make(map[string]interface{}, 0)
	itemMap["episode_id"] = video.EpisodeID
	itemMap["section_id"] = video.SectionID
	itemMap["name"] = video.Name
	itemMap["keywords"] = video.Keywords
	itemMap["description"] = video.Description
	itemMap["cover_key"] = video.CoverKey
	itemMap["mp4_key"] = video.Mp4Key
	itemMap["duration"] = video.Duration
	itemMap["is_publish"] = video.IsPublish
	itemMap["is_free"] = video.IsFree

	if err := srv.UpdateVideo(itemMap, videoId); err != nil {
		app.Response(c, errno.InternalServerError, nil)
		return
	}

	app.Response(c, nil, nil)
}

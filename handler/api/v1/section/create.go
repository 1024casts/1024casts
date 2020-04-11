package section

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
// @Router /sections [post]
func Create(c *gin.Context) {
	log.Info("Video Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	item := model.SectionModel{
		CourseId: r.CourseId,
		Name:     r.Name,
		Weight:   r.Weight,
	}

	srv := service.NewCourseService()
	id, err := srv.CreateSection(item)
	if err != nil {
		log.Warn("section info", lager.Data{"id": id})
		app.Response(c, errno.ErrSectionCreateFail, nil)
		return
	}

	resp := CreateResponse{
		Id: item.Id,
	}

	// Show the user information.
	app.Response(c, nil, resp)
}

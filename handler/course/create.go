package course

import (
	. "1024casts/backend/handler"
	"1024casts/backend/model"
	"1024casts/backend/pkg/errno"
	"1024casts/backend/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user [post]
func Create(c *gin.Context) {
	log.Info("Course Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	course := model.CourseModel{
		Name: r.Name,
		Type: r.Type,
		Description:    r.Description,
		Slug:r.Slug,
		CoverImage: r.CoverImage,
		IsPublish: r.IsPublish,
	}

	// Insert the user to the database.
	if err := course.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	resp := CreateResponse{
		Name: course.Name,
	}

	// Show the user information.
	SendResponse(c, nil, resp)
}

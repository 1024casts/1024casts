package course

import (
	"strconv"

	. "1024casts/backend/handler"
	"1024casts/backend/model"
	"1024casts/backend/pkg/errno"
	"1024casts/backend/util"

	"1024casts/backend/service"

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
// @Router /courses/{id} [put]
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// Get the course id from the url parameter.
	courseId, _ := strconv.Atoi(c.Param("id"))

	// Binding the course data.
	var course model.CourseModel
	if err := c.Bind(&course); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// We update the record based on the user id.
	course.Id = uint64(courseId)

	srv := service.NewCourseService()
	//user, err := model.GetUserById(userId)
	_, err := srv.GetCourseById(courseId)
	if err != nil {
		SendResponse(c, errno.ErrCourseNotFound, nil)
		log.Warn("course info", lager.Data{"id": courseId})
		return
	}

	// Validate the data.
	//if err := u.Validate(); err != nil {
	//	SendResponse(c, errno.ErrValidation, nil)
	//	return
	//}

	// Save changed fields.
	courseMap := make(map[string]interface{}, 0)
	courseMap["name"] = course.Name
	courseMap["type"] = course.Type
	courseMap["description"] = course.Description
	courseMap["slug"] = course.Slug
	courseMap["cover_image"] = course.CoverImage
	courseMap["is_publish"] = course.IsPublish
	courseMap["update_status"] = course.UpdateStatus

	if err := srv.UpdateCourse(courseMap, courseId); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}

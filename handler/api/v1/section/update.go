package section

import (
	"strconv"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/errno"
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
// @Router /sections/{id} [put]
func Update(c *gin.Context) {
	log.Info("Section Update function called.")
	// Get the course id from the url parameter.
	sectionId, _ := strconv.Atoi(c.Param("id"))

	// Binding the course data.
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	srv := service.NewCourseService()
	_, err := srv.GetSectionById(sectionId)
	if err != nil {
		log.Warn("section info", lager.Data{"id": sectionId})
		app.Response(c, errno.ErrVideoNotFound, nil)
		return
	}

	// Save changed fields.
	itemMap := make(map[string]interface{}, 0)
	itemMap["course_id"] = r.CourseId
	itemMap["name"] = r.Name
	itemMap["weight"] = r.Weight

	if err := srv.UpdateSection(itemMap, sectionId); err != nil {
		app.Response(c, errno.InternalServerError, nil)
		return
	}

	app.Response(c, nil, nil)
}

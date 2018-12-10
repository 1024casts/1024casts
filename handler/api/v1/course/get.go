package course

import (
	"strconv"

	. "github.com/1024casts/1024casts/handler"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary Get a course by the course identifier
// @Description Get a course by id
// @Tags course
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} model.CourseModel "{"code":0,"message":"OK","data":{"username":"kong","password":"$2a$10$E0kwtmtLZbwW/bDQ8qI8e.eHPqhQOW9tvjwpyo/p05f/f4Qvr3OmS"}}"
// @Router /courses/{id} [get]
func Get(c *gin.Context) {
	log.Info("Get function called.")

	// Get the user by the `username` from the database.
	//username := c.Param("username")
	//user, err := model.GetUser(username)

	// Get the user by the `id` from the database.
	// Get the user id from the url parameter.
	courseId, _ := strconv.Atoi(c.Param("id"))

	srv := service.NewCourseService()
	//user, err := model.GetUserById(userId)
	course, err := srv.GetCourseById(courseId)
	if err != nil {
		SendResponse(c, errno.ErrCourseNotFound, nil)
		log.Warn("course info", lager.Data{"id": courseId})
		return
	}

	SendResponse(c, nil, course)
}

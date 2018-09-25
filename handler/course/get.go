package course

import (
	"strconv"

	. "1024casts/backend/handler"
	"1024casts/backend/model"
	"1024casts/backend/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
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

	course, err := model.GetCourseById(courseId)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, course)
}

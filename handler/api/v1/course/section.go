package course

import (
	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/service"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// @Summary List the courses in the database
// @Description List courses
// @Tags course
// @Accept  json
// @Produce  json
// @Param course body course.ListRequest true "List courses"
// @Success 200 {object} course.SwaggerListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /courses [get]
func Section(c *gin.Context) {
	log.Info("Section function called.")

	courseId, _ := strconv.Atoi(c.Param("id"))

	srv := service.NewCourseService()
	infos, err := srv.GetCourseSectionListWithVideo(uint64(courseId))
	if err != nil {
		app.Response(c, err, nil)
		return
	}

	app.Response(c, nil, SectionListResponse{
		TotalCount: 100,
		List:       infos,
	})
}

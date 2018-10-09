package course

import (
	. "1024casts/backend/handler"
	"1024casts/backend/service"

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
	courseIdd := uint64(courseId)

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("get page error", err)
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		log.Error("get limit error", err)
	}

	offset := (page - 1) * limit

	srv := service.NewCourseService()
	infos, count, err := srv.GetCourseSectionList(courseIdd, offset, limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, SectionListResponse{
		TotalCount: count,
		List:       infos,
	})
}

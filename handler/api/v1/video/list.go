package video

import (
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"

	"strconv"
	"strings"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// @Summary List the videos in the database
// @Description List videos
// @Tags video
// @Accept  json
// @Produce  json
// @Param video body video.ListRequest true "List videos"
// @Success 200 {object} video.SwaggerListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /videos [get]
func List(c *gin.Context) {
	log.Info("List function called.")
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	courseId, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		log.Error("get course_id error", err)
	}

	name := strings.TrimSpace(c.Query("name"))

	srv := service.NewVideoService()
	infos, count, err := srv.GetVideoListPagination(uint64(courseId), name, 0, 100)
	if err != nil {
		app.Response(c, err, nil)
		return
	}

	app.Response(c, nil, ListResponse{
		TotalCount: count,
		List:       infos,
	})
}

package video

import (
	. "1024casts/backend/handler"
	"1024casts/backend/pkg/errno"
	"1024casts/backend/service"
	"strconv"
	"strings"

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
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	courseId, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		log.Error("get course_id error", err)
	}

	name := strings.TrimSpace(c.Query("name"))

	srv := service.NewVideoService()
	infos, count, err := srv.GetVideoList(uint64(courseId), name, 0, 100)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		List:       infos,
	})
}

package course

import (
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"

	"strconv"

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
func Video(c *gin.Context) {
	log.Info("Video function called.")
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error("get course_id error", err)
	}

	srv := service.NewVideoService()
	infos, err := srv.GetVideoList(uint64(courseId), false)
	if err != nil {
		app.Response(c, err, nil)
		return
	}

	app.Response(c, nil, VideoListResponse{
		TotalCount: 0,
		List:       infos,
	})
}

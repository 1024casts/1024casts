package qiniu

import (
	"fmt"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Upload(c *gin.Context) {

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		log.Warnf("[upload] get file err: %v", err)
		app.Response(c, errno.ErrGetUploadFile, nil)
		return
	}

	qiNiuSrv := service.NewQiNiuService()
	uploadRet, err := qiNiuSrv.UploadImage(c, file)
	if err != nil {
		log.Warnf("[upload] upload file err: %v", err)
		app.Response(c, errno.ErrUploadingFile, nil)
		return
	}
	fmt.Println(uploadRet)

	app.Response(c, nil, uploadRet)

}

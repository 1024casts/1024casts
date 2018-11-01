package qiniu

import (
	"context"
	"fmt"

	. "1024casts/backend/handler"

	"1024casts/backend/util"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/spf13/viper"
)

func Upload(c *gin.Context) {

	accessKey := viper.GetString("qiniu.AccessKey")
	secretKey := viper.GetString("qiniu.SecretKey")
	bucket := viper.GetString("qiniu.Bucket")

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		log.Warn("[qiniu] get file err", lager.Data{"error": err})
		return
	}

	saveRootPath := viper.GetString("upload.dst")
	imagePrefix := "uploads/images/" + util.GetDate() + "/"
	imagePath := saveRootPath + imagePrefix
	if err := os.MkdirAll(imagePath, 0777); err != nil {
		log.Fatal("[qiniu] create dir err", err)
		return
	}

	key := imagePrefix + file.Filename

	// Upload the file to specific dst.
	dst := saveRootPath + key
	if err := c.SaveUploadedFile(file, dst); err != nil {
		log.Fatal("[qiniu] upload file err", err)
		return
	}

	localFile := dst
	putPolicy := storage.PutPolicy{
		Scope: bucket + ":" + key,
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": key,
		},
	}
	//putExtra.NoCrc32Check = true
	if err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)

	SendResponse(c, nil, UploadResponse{
		Key:  ret.Key,
		Hash: ret.Hash,
	})

}

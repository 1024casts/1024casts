package service

import (
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/repository"

	"github.com/1024casts/1024casts/pkg/constvar"

	"fmt"
	"os"

	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/spf13/viper"
)

type QiNiuService struct {
}

func NewQiNiuService() *QiNiuService {
	return &QiNiuService{}
}

type UploadResponse struct {
	Key  string `json:"key"`
	Hash string `json:"hash"`
	Url  string `json:"url"`
}

func (srv *QiNiuService) UploadImage(c *gin.Context, file *multipart.FileHeader, isPublicBucket bool) (resp UploadResponse, err error) {
	saveRootPath := viper.GetString("upload.dst")
	imagePrefix := "uploads/images/" + util.GetDate() + "/"
	imagePath := saveRootPath + imagePrefix
	if err = os.MkdirAll(imagePath, 0777); err != nil {
		log.Fatal("[qiniu] create dir err", err)
		return
	}

	fileNameWithoutExt, err := util.GenShortId()
	if err != nil {
		log.Warnf("[qiniu] gen filename err, %v", err)
		return
	}
	filename := fileNameWithoutExt + filepath.Ext(file.Filename)
	key := imagePrefix + filename

	// Upload the file to specific dst.
	dst := saveRootPath + key
	if err = c.SaveUploadedFile(file, dst); err != nil {
		log.Fatal("[qiniu] upload file err", err)
		return
	}
	localFile := dst
	log.Infof("[qiniu] local_file: %s", localFile)

	accessKey := viper.GetString("qiniu.access_key")
	secretKey := viper.GetString("qiniu.secret_key")
	bucket := viper.GetString("qiniu.bucket")
	if isPublicBucket {
		bucket = viper.GetString("qiniu.public_bucket")
	}

	putPolicy := storage.PutPolicy{
		Scope: bucket,
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
			"x:name": filename,
		},
	}
	//putExtra.NoCrc32Check = true
	if err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra); err != nil {
		fmt.Println(err)
		return
	}

	log.Infof("[qiniu] uploaded file ret: %v", ret)

	resp.Key = ret.Key
	resp.Hash = ret.Hash
	if isPublicBucket {
		resp.Url = util.GetQiNiuPublicAccessUrl(ret.Key)
	} else {
		resp.Url = util.GetQiNiuPrivateAccessUrl(ret.Key, constvar.MediaTypeImage, 300, 0)
	}

	// write to image table
	imageRepo := repository.NewImageRepo()
	imgModel := model.ImageModel{}
	imgModel.UserID = util.GetUserId(c)
	imgModel.ImageName = fileNameWithoutExt
	imgModel.ImagePath = ret.Key
	_, err = imageRepo.CreateImage(imgModel)
	if err != nil {
		log.Warnf("[qiniu] add to image table err, %v", err)
	}

	return resp, nil

}

func (srv *QiNiuService) UploadVideo(c *gin.Context, file *multipart.FileHeader, isPublicBucket bool) (resp UploadResponse, err error) {
	saveRootPath := viper.GetString("upload.dst")
	imagePrefix := "uploads/video/" + util.GetDate() + "/"
	imagePath := saveRootPath + imagePrefix
	if err = os.MkdirAll(imagePath, 0777); err != nil {
		log.Fatal("[qiniu] create dir err", err)
		return
	}

	fileNameWithoutExt, err := util.GenShortId()
	if err != nil {
		log.Warnf("[qiniu] gen filename err, %v", err)
		return
	}
	filename := fileNameWithoutExt + filepath.Ext(file.Filename)
	key := imagePrefix + filename

	// Upload the file to specific dst.
	dst := saveRootPath + key
	if err = c.SaveUploadedFile(file, dst); err != nil {
		log.Fatal("[qiniu] upload file err", err)
		return
	}
	localFile := dst
	log.Infof("[qiniu] local_file: %s", localFile)

	accessKey := viper.GetString("qiniu.access_key")
	secretKey := viper.GetString("qiniu.secret_key")
	bucket := viper.GetString("qiniu.bucket")
	if isPublicBucket {
		bucket = viper.GetString("qiniu.public_bucket")
	}

	putPolicy := storage.PutPolicy{
		Scope: bucket,
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
			"x:name": filename,
		},
	}
	//putExtra.NoCrc32Check = true
	if err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra); err != nil {
		fmt.Println(err)
		return
	}

	log.Infof("[qiniu] uploaded file ret: %v", ret)

	resp.Key = ret.Key
	resp.Hash = ret.Hash
	if isPublicBucket {
		resp.Url = util.GetQiNiuPublicAccessUrl(ret.Key)
	} else {
		resp.Url = util.GetQiNiuPrivateAccessUrl(ret.Key, constvar.MediaTypeVideo, 300, 0)
	}

	// write to image table
	imageRepo := repository.NewImageRepo()
	imgModel := model.ImageModel{}
	imgModel.UserID = util.GetUserId(c)
	imgModel.ImageName = fileNameWithoutExt
	imgModel.ImagePath = ret.Key
	_, err = imageRepo.CreateImage(imgModel)
	if err != nil {
		log.Warnf("[qiniu] add to image table err, %v", err)
	}

	return resp, nil

}

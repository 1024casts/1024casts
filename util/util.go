package util

import (
	"time"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/spf13/viper"
	"github.com/teris-io/shortid"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

// 获取七牛资源的私有链接
func GetQiniuPrivateAccessUrl(path string) string {
	accessKey := viper.GetString("qiniu.AccessKey")
	secretKey := viper.GetString("qiniu.SecretKey")
	mac := qbox.NewMac(accessKey, secretKey)

	domain := viper.GetString("qiniu.CDN_URL")
	key := strings.TrimPrefix(path, "/")

	imageStyle := "imageView2/2/w/200/h/200/q/75|imageslim"
	key = key + "?" + imageStyle
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期

	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)

	return privateAccessURL
}

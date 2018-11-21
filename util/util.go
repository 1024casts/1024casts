package util

import (
	"time"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	hashids "github.com/speps/go-hashids"
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

// encode uid 为字符串
func EncodeUid(uid int64) string {
	hd := hashids.NewData()
	hd.Salt = "1024casts_uid"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	e, _ := h.EncodeInt64([]int64{uid})
	return e

}

// decode uid 为int64
func DecodeUid(encodedUid string) (uid int64) {
	hd := hashids.NewData()
	hd.Salt = "1024casts_uid"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	d, _ := h.DecodeInt64WithError(encodedUid)
	if len(d) > 0 {
		return d[0]
	}

	return 0

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

func TimestampToString(ts time.Time) string {
	return time.Unix(ts.Unix(), 00).Format("2006-01-02 15:04:05")
}

func GetDate() string {
	return time.Now().Format("2006/01/02")
}

package util

import (
	"time"

	"strings"

	"fmt"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
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
	e, err := h.EncodeInt64([]int64{uid})
	if err != nil {
		log.Warn("encode uid err")
	}

	return e

}

// decode uid 为int64
func DecodeUid(encodedUid string) (uid int64) {
	hd := hashids.NewData()
	hd.Salt = "1024casts_uid"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	d, err := h.DecodeInt64WithError(encodedUid)

	if err != nil {
		log.Warn("decode uid err")
	}

	if len(d) > 0 {
		return d[0]
	}

	return 0

}

func GetUserId(ctx *gin.Context) uint64 {
	cookie, err := ctx.Cookie(viper.GetString("cookie_name"))
	if err != nil {
		log.Warnf("[util] get cookie err: %+v", err)
		return 0
	}

	userId := DecodeUid(cookie)
	return uint64(userId)
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

func GetImageFullUrl(uri string) string {
	return viper.GetString("image_domain") + uri
}

func GenerateOrderNo() (uint64, error) {
	dateStr := time.Now().Format("20060102150405")
	log.Infof("data str: %s", dateStr)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randStr := fmt.Sprintf("%05v", rnd.Intn(10000))
	log.Infof("rand str: %s", randStr)

	orderNoStr := dateStr + randStr
	orderNo, err := strconv.ParseUint(orderNoStr, 10, 64)
	if err != nil {
		log.Warnf("[util] convert: %s err: %+v", randStr, err)
		return 0, err
	}

	log.Infof("orderNo: %d", orderNo)

	return orderNo, nil
}
